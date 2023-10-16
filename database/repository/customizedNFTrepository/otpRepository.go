package customizedNFTrepository

//*OTP -> One Time Password
import (
	"context"
	"fmt"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/dtos/responseDtos"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserAuth = "userAuth"
var walletnft = "walletnft"
var walletTenant = "walletnfttenant"

type OtpRepository struct{}

/**
 * Description : Save new OTP along with email and batch ID in ruriOtp Collection
 **Params : otpDataSet, OTPData struct containting data to be stored.
 **reutrns : objectID if dat gets stored or an error if it dosnt
 */
func (r *OtpRepository) SaveOTP(otpDataSet models.UserAuth) (string, error) {
	var authrst models.UserAuth
	rst, err := repository.FindById1AndId2("email", otpDataSet.Email, "shopid", otpDataSet.ShopID, UserAuth)
	for rst.Next(context.TODO()) {
		err = rst.Decode(&authrst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return "", err
		}
	}
	if authrst.Email != "" {
		if authrst.Validated {
			err := fmt.Errorf("OTP already validated")
			return "", err
		}
		err := fmt.Errorf("OTP for this email already exists")
		return "", err
	}
	return repository.Save(otpDataSet, UserAuth)
}

/**
 * Descprition : checks if a valid OTP exisit in collcetion ruriOtp
 * *param : email, users email
 * *param : otp, otp entered by user
 * *reutrns : respective batchID if the otp is valid
 */
func (r *OtpRepository) ValidateOTP(email string, otp string) (string, string, error) {
	var authrst models.UserAuth
	rst, err := repository.FindById1AndId2("email", email, "otp", otp, UserAuth)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&authrst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return "", "", err
		}
	}

	if authrst.BatchID == "" {
		return "Invalid OTP", "", err

	} else {
		rest, err := UpdateOTPStatus(authrst.OtpID)
		if !rest {
			logs.ErrorLogger.Println("Failed to update DB: ", err.Error())
			return "Update Failed", "", err
		}
	}
	now := primitive.NewDateTimeFromTime(time.Now())
	// Below commented out code can be used to verify the expiration date check
	// duration := time.Hour * 24 * 60
	// Dummydate := primitive.NewDateTimeFromTime(Dummydate.Add(duration))
	// logs.InfoLogger.Println("data from DB:", authrst)
	// logs.InfoLogger.Println("NOW: ", Dummydate)
	// logs.InfoLogger.Println("EXP: ", authrst.ExpireDate)
	if now > authrst.ExpireDate {
		return "Expired OTP", "", nil
	}
	return authrst.BatchID, authrst.ShopID, nil
}
func UpdateOTPStatus(otpID primitive.ObjectID) (bool, error) {
	update := bson.M{
		"$set": bson.M{"validated": true},
	}
	session, err := connections.GetMongoSession()
	if err != nil {
		logs.ErrorLogger.Println("Error while getting session " + err.Error())
	}
	defer session.EndSession(context.TODO())
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	rst := session.Client().Database(connections.DbName).Collection(UserAuth).FindOneAndUpdate(context.TODO(), bson.M{"_id": otpID}, update, &opt)
	if rst != nil {
		err := rst.Decode(&rst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreiving data from DB : ", err.Error())
			return false, err
		}
	}
	return true, nil
}

func (r *OtpRepository) ResendOTP(otpDataSet models.UserAuth) (string, error) {
	var authrst models.UserAuth
	rst, err := repository.FindById1AndId2("email", otpDataSet.Email, "batchid", otpDataSet.BatchID, UserAuth)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&authrst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return "", err
		}
	}
	if authrst.Validated {
		err := fmt.Errorf("OTP already validated")
		return "", err
	}
	if authrst.BatchID == "" { //*IF OTP was not stored in DB new entry will be made
		return repository.Save(otpDataSet, UserAuth)
	} else { //* IF OTP data already exisit it will get updated
		update := bson.M{
			"$set": bson.M{"otp": otpDataSet.Otp, "expDate": otpDataSet.ExpireDate},
		}
		session, err := connections.GetMongoSession()
		if err != nil {
			logs.ErrorLogger.Println("Error while getting session " + err.Error())
		}

		defer session.EndSession(context.TODO())
		upsert := false
		after := options.After
		opt := options.FindOneAndUpdateOptions{
			ReturnDocument: &after,
			Upsert:         &upsert,
		}
		rst := session.Client().Database(connections.DbName).Collection(UserAuth).FindOneAndUpdate(context.TODO(), bson.M{"email": otpDataSet.Email, "batchid": otpDataSet.BatchID}, update, &opt)
		var responseOtp models.UserAuth
		if rst != nil {
			err := rst.Decode(&responseOtp)
			if err != nil {
				logs.InfoLogger.Println("Failed to update DB: ", err.Error())
				return responseOtp.BatchID, err
			} else {
				return responseOtp.BatchID, err
			}
		} else {
			return responseOtp.BatchID, err
		}
	}
}

func (r *OtpRepository) ValidateNFTStatus(email string, otp string) (string, error) {
	var authrst models.WalletNFT
	rst, err := repository.FindById1AndId2("email", email, "otp", otp, walletnft)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&authrst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return authrst.NFTStatus, err
		}
	}
	return authrst.NFTStatus, nil
}

func (r *OtpRepository) ValidateNFTStatusbyShopId(shopid string) (string, error) {
	var authrst models.WalletNFT
	rst, err := repository.FindById("shopid", shopid, walletnft)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&authrst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retrieving data from collection walletnft in ValidateOTP:OtpRepository.go: ", err.Error())
			return authrst.NFTStatus, err
		}
	}
	return authrst.NFTStatus, nil
}

func (r *OtpRepository) GetWalletTenant(name string) (models.WalletNFTTenantUser, error) {
	var tenant models.WalletNFTTenantUser
	rst, err := repository.FindById("name", name, walletTenant)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	for rst.Next(context.TODO()) {
		decodeErr := rst.Decode(&tenant)
		if decodeErr != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return tenant, err
		}
	}
	if tenant.Name != "" {
		return tenant, nil
	} else {
		err := fmt.Errorf("Tenant not found")
		return tenant, err
	}
}
func (r *OtpRepository) CheckOTPValidatedStatus(email string, id string) (responseDtos.OTPStatus, error) {
	var authrst models.UserAuth
	var response responseDtos.OTPStatus
	rst, err := repository.FindById1AndId2("email", email, "shopid", id, UserAuth)
	for rst.Next(context.TODO()) {
		err = rst.Decode(&authrst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection in ValidateOTP:OtpRepository.go: ", err.Error())
			return response, err
		}
	}
	if authrst.Email != "" {
		if authrst.Validated {
			err := fmt.Errorf("OTP already validated")
			response.Message = err.Error()
			response.IsOTPValidated = true
			return response, nil
		}
		err := fmt.Errorf("OTP Sent but not validated")
		response.Message = err.Error()
		response.IsOTPValidated = false
		return response, nil
	}
	noOTPError := fmt.Errorf("OTP does not exist for this email")
	return response, noOTPError
}
