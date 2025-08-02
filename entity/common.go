package entity

import (
	"time"

	"github.com/Amierza/TedXBackend/constants"
	"gorm.io/gorm"
)

type (
	TimeStamp struct {
		CreatedAt time.Time      `gorm:"column:createdAt" json:"created_at"`
		UpdatedAt time.Time      `gorm:"column:updatedAt" json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"column:deletedAt" json:"deleted_at"`
	}

	Role                string
	AudienceType        string
	SponsorshipCategory string
	MerchCategory       string
	Instansi            string
	ItemType            string
	BundleType          string
	TicketType          string
)

const (
	Admin Role = constants.ENUM_ROLE_ADMIN
	Guest Role = constants.ENUM_ROLE_GUEST

	PreEvent3 TicketType = constants.ENUM_TICKET_PRE_EVENT_3
	MainEvent TicketType = constants.ENUM_TICKET_MAIN_EVENT

	Sponsor      SponsorshipCategory = constants.ENUM_SPONSORSHIP_CATEGORY_SPONSOR
	Partner      SponsorshipCategory = constants.ENUM_SPONSORSHIP_CATEGORY_PARTNER
	MediaPartner SponsorshipCategory = constants.ENUM_SPONSORSHIP_CATEGORY_MEDIA_PARTNER

	TShirt  MerchCategory = constants.ENUM_MERCH_CATEGORY_TSHIRT
	Cap     MerchCategory = constants.ENUM_MERCH_CATEGORY_CAP
	Sticker MerchCategory = constants.ENUM_MERCH_CATEGORY_STICKER
	Other   MerchCategory = constants.ENUM_MERCH_CATEGORY_OTHER

	Unair Instansi = constants.ENUM_INSTANSI_UNAIR
	Umum  Instansi = constants.ENUM_INSTANSI_UMUM

	TicketItemType ItemType = constants.ENUM_TICKET_ITEM_TYPE
	MerchItemType  ItemType = constants.ENUM_MERCH_ITEM_TYPE
	BundleItemType ItemType = constants.ENUM_BUNDLE_ITEM_TYPE

	BundleMerchType       BundleType = constants.ENUM_BUNDLE_MERCH_TYPE
	BundleMerchTicketType BundleType = constants.ENUM_BUNDLE_MERCH_TICKET_TYPE

	Regular AudienceType = constants.ENUM_AUDIENCE_REGULAR
	Invited AudienceType = constants.ENUM_AUDIENCE_INVITED
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

func IsValidInstansi(i Instansi) bool {
	return i == Unair || i == Umum
}

func IsValidTicketType(t TicketType) bool {
	return t == PreEvent3 || t == MainEvent
}

func IsValidItemType(it ItemType) bool {
	return it == TicketItemType || it == MerchItemType || it == BundleItemType
}

func IsValidBundleType(bt BundleType) bool {
	return bt == BundleMerchTicketType || bt == BundleMerchType
}
