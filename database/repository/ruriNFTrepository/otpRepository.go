package ruriNFTrepository

import (
	"context"
	"fmt"

	"github.com/dileepaj/tracified-nft-backend/database/connections"
	"github.com/dileepaj/tracified-nft-backend/database/repository"
	"github.com/dileepaj/tracified-nft-backend/models"
	"github.com/dileepaj/tracified-nft-backend/utilities/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RurriOTP = "ruriOtp"

type OtpRepository struct{}

/**
 * Description : Save new OTP along with email and batch ID in ruriOtp Collection
 **Params : otpDataSet, OTPData struct containting data to be stored.
 **reutrns : objectID if dat gets stored or an error if it dosnt
 */
func (r *Rurirepository) SaveOTP(otpDataSet models.OTPData) (string, error) {
	return repository.Save(otpDataSet, RurriOTP)
}

/**
 * Descprition : checks if a valid OTP exisit in collcetion ruriOtp
 * *param : email, users email
 * *param : otp, otp entered by user
 * *reutrns : respective batchID if the otp is valid
 */
func (r *Rurirepository) ValidateOTP(email string, otp string) (string, error) {
	var otprst models.OTPData
	rst, err := repository.FindById1AndId2("email", email, "otp", otp, RurriOTP)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&otprst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return "", err
		}
	}
	logs.InfoLogger.Println("data retrived from DB :", otprst)
	if otprst.BatchID == "" {
		return "Invalid OTP", err
	} else {
		return otprst.BatchID, err
	}

}
func (r *Rurirepository) ResendOTP(otpDataSet models.OTPData) (string, error) {
	var otprst models.OTPData
	rst, err := repository.FindById1AndId2("email", otpDataSet.Email, "batchid", otpDataSet.BatchID, RurriOTP)
	if err != nil {
		logs.ErrorLogger.Println("failed to return data from DB: ", err.Error())
	}
	for rst.Next(context.TODO()) {
		err = rst.Decode(&otprst)
		if err != nil {
			logs.ErrorLogger.Println("Error occured while retreving data from collection ruriotp in ValidateOTP:OtpRepository.go: ", err.Error())
			return "", err
		}
	}
	if otprst.BatchID == "" { //*IF OTP was not stored in DB new entry will be made
		fmt.Println("data not recorded in DB")
		return repository.Save(otpDataSet, RurriOTP)
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
		rst := session.Client().Database(connections.DbName).Collection("ruriOtp").FindOneAndUpdate(context.TODO(), bson.M{"email": otpDataSet.Email}, update, &opt)
		var responseOtp models.OTPData
		if rst != nil {
			err := rst.Decode(&responseOtp)
			if err != nil {
				logs.InfoLogger.Println("Failed toupdate DB")
				return responseOtp.BatchID, err
			} else {
				return responseOtp.BatchID, err
			}
		} else {
			return responseOtp.BatchID, err
		}
	}
}
