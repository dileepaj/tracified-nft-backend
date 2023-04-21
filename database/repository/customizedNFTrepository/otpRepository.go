package customizedNFTrepository

//*OTP -> One Time Password
import (
	"context"
	"fmt"
	"time"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserAuth = "userAuth"
var RURI = "rurinft"

type OtpRepository struct{}

/**
 * Description : Save new OTP along with email and batch ID in ruriOtp Collection
 **Params : otpDataSet, OTPData struct containting data to be stored.
 **reutrns : objectID if dat gets stored or an error if it dosnt
 */
func (r *OtpRepository) SaveOTP(otpDataSet models.UserAuth) (string, error) {
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
		"$set": bson.M{"validated": "True"},
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
	if authrst.BatchID == "" { //*IF OTP was not stored in DB new entry will be made
		fmt.Println("data not recorded in DB")
		return repository.Save(otpDataSet, UserAuth)
	} else { //* IF OTP data already exisit it will get updated
		fmt.Println("data was recorded in DB Updating")
		update := bson.M{
			"$set": bson.M{"otp": otpDataSet.Otp},
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
		rst := session.Client().Database(connections.DbName).Collection(UserAuth).FindOneAndUpdate(context.TODO(), bson.M{"email": otpDataSet.Email}, update, &opt)
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
	rst, err := repository.FindById1AndId2("email", email, "otp", otp, RURI)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	logs.InfoLogger.Println("rst : ", rst)
	for rst.Next(context.TODO()) {
		err = rst.Decode(&authrst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return authrst.NFTStatus, err
		}
	}
	return authrst.NFTStatus, nil
}
