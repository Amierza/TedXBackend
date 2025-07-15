package entity

import (
	"time"

	"github.com/Amierza/TedXBackend/constants"
	"gorm.io/gorm"
)

type (
	TimeStamp struct {
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`
	}

	Role                string
	AudienceType        string
	SponsorshipCategory string
	MerchCategory       string
	TransactionStatus   string
	PaymentType         string
	Acquire             string
	ItemType            string
	BundleType          string
)

const (
	Admin Role = constants.ENUM_ROLE_ADMIN
	Guest Role = constants.ENUM_ROLE_GUEST

	Regular AudienceType = constants.ENUM_AUDIENCE_REGULAR
	Invited AudienceType = constants.ENUM_AUDIENCE_INVITED

	Sponsor      SponsorshipCategory = constants.ENUM_SPONSORSHIP_CATEGORY_SPONSOR
	Partner      SponsorshipCategory = constants.ENUM_SPONSORSHIP_CATEGORY_PARTNER
	MediaPartner SponsorshipCategory = constants.ENUM_SPONSORSHIP_CATEGORY_MEDIA_PARTNER

	TShirt  MerchCategory = constants.ENUM_MERCH_CATEGORY_TSHIRT
	Cap     MerchCategory = constants.ENUM_MERCH_CATEGORY_CAP
	Sticker MerchCategory = constants.ENUM_MERCH_CATEGORY_STICKER
	Other   MerchCategory = constants.ENUM_MERCH_CATEGORY_OTHER

	StatusPending    TransactionStatus = constants.ENUM_MIDTRANS_STATUS_PENDING
	StatusSettlement TransactionStatus = constants.ENUM_MIDTRANS_STATUS_SETTLEMENT
	StatusFailure    TransactionStatus = constants.ENUM_MIDTRANS_STATUS_FAILURE
	StatusExpired    TransactionStatus = constants.ENUM_MIDTRANS_STATUS_EXPIRE
	StatusCancelled  TransactionStatus = constants.ENUM_MIDTRANS_STATUS_CANCEL
	StatusDenied     TransactionStatus = constants.ENUM_MIDTRANS_STATUS_DENY
	StatusRefund     TransactionStatus = constants.ENUM_MIDTRANS_STATUS_REFUND

	PaymentTypeBankTransfer PaymentType = constants.ENUM_MIDTRANS_PAYMENT_TYPE_BANK_TRANSFER
	PaymentTypeCreditCard   PaymentType = constants.ENUM_MIDTRANS_PAYMENT_TYPE_CREDIT_CARD
	PaymentTypeQRIS         PaymentType = constants.ENUM_MIDTRANS_PAYMENT_TYPE_QRIS
	PaymentTypeGopay        PaymentType = constants.ENUM_MIDTRANS_PAYMENT_TYPE_GOPAY
	PaymentTypeShopeePay    PaymentType = constants.ENUM_MIDTRANS_PAYMENT_TYPE_SHOPEE

	AcquireMandiri Acquire = constants.ENUM_MIDTRANS_ACQUIRE_MANDIRI
	AcquireBRI     Acquire = constants.ENUM_MIDTRANS_ACQUIRE_BRI
	AcquireBCA     Acquire = constants.ENUM_MIDTRANS_ACQUIRE_BCA
	AcquireBNI     Acquire = constants.ENUM_MIDTRANS_ACQUIRE_BNI
	AcquireCIMB    Acquire = constants.ENUM_MIDTRANS_ACQUIRE_CIMB

	TicketItemType ItemType = constants.ENUM_TICKET_ITEM_TYPE
	MerchItemType  ItemType = constants.ENUM_MERCH_ITEM_TYPE
	BundleItemType ItemType = constants.ENUM_BUNDLE_ITEM_TYPE

	BundleMerchType       BundleType = constants.ENUM_BUNDLE_MERCH_TYPE
	BundleMerchTicketType BundleType = constants.ENUM_BUNDLE_MERCH_TICKET_TYPE
)

func IsValidRole(r Role) bool {
	return r == Admin || r == Guest
}

func IsValidAudienceType(at AudienceType) bool {
	return at == Regular || at == Invited
}

func IsValidSponsorshipCategory(sc SponsorshipCategory) bool {
	return sc == Sponsor || sc == Partner || sc == MediaPartner
}

func IsValidMerchCategory(mc MerchCategory) bool {
	return mc == TShirt || mc == Cap || mc == Sticker || mc == Other
}

func IsValidTransactionStatus(ts TransactionStatus) bool {
	return ts == StatusPending || ts == StatusSettlement || ts == StatusFailure || ts == StatusExpired || ts == StatusCancelled || ts == StatusDenied || ts == StatusRefund
}

func IsValidPaymentType(pt PaymentType) bool {
	return pt == PaymentTypeBankTransfer || pt == PaymentTypeCreditCard || pt == PaymentTypeQRIS || pt == PaymentTypeGopay || pt == PaymentTypeShopeePay
}

func IsValidAcquire(a Acquire) bool {
	return a == AcquireMandiri || a == AcquireBRI || a == AcquireBCA || a == AcquireBNI || a == AcquireCIMB
}

func IsValidItemType(it ItemType) bool {
	return it == TicketItemType || it == MerchItemType || it == BundleItemType
}

func IsValidBundleType(bt BundleType) bool {
	return bt == BundleMerchTicketType || bt == BundleMerchType
}
