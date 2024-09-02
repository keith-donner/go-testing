package o

import "time"

type Office struct {
	ID       string
	CityCode string
}

type SourceType int

const (
	GDS SourceType = iota
	NDC
	DIRECT
)

func (s SourceType) String() string {
	return [...]string{"GDS", "NDC", "DIRECT"}[s]
}

type Source struct {
	System string
	Type   SourceType
}

type SeatPosition int

const (
	Left SeatPosition = iota
	Forward
	Aisle
	Window
	Exit
	Bulkhead
)

func (sp SeatPosition) String() string {
	return [...]string{"Left", "Forward", "Aisle", "Window", "Exit", "Bulkhead"}[sp]
}

type Seat struct {
	Number        string
	Position      SeatPosition
	SegmentNumber int
	Status        string
	Source        Source
}

type Queue struct {
	Name          string
	Number        int
	Category      string
	SabrePrefCode string
	Source        Source
}

type Name struct {
	First  string
	Middle string
	Last   string
	Prefix string
	Suffix string
}

type NameNumber struct {
	Prefix string
	Suffix string
	String string
}

type SpecialMeal struct {
	MealCode       string
	Status         string
	SegmentNumbers []int
}

type FrequentFlyer struct {
	AssocAirline   string
	Number         string
	AirlineStatus  string
	FqtvAirline    string
	SegmentNumbers []int
}

type Card struct {
	Code        string
	Number      string
	ExpiryMonth int
	ExpiryYear  int
}

type Country struct {
	Name string
	Code string
}

type StateProvince struct {
	Name string
	Code string
}

type ExchangeItem struct {
	Adcol           float64
	ChangeFee       float64
	Document        string
	Forfeit         float64
	IssueDate       time.Time
	NewTicketTotal  float64
	OrigName        string
	OrigTicketTotal float64
	Payment         float64
	Residual        float64
}

type InvoiceItem struct {
	Description string
	Payment     Payment
	State       string
	Name        string
	Tax         float64
	Document    string
	Type        string
	Base        float64
	Total       float64
	Exchange    ExchangeItem
}

type Distance struct {
	Miles float64
	Kms   float64
}
type FareTax struct {
	Code        string
	Amount      float64
	Currency    string
	IncludesVAT bool
}

type FareRestriction struct {
	IsRefundable         bool
	CancelPenaltyPercent float64
	ChangePenaltyDollar  float64
}

type FareSegmentExtra struct {
	FareBasis        string
	TicketDesignator string
	NotValidBefore   time.Time
	NotValidAfter    time.Time
	SegmentNumber    int
}

type FareSegment struct {
	ClassOfService       string
	DepartureAirportCode string
	ArrivalAirportCode   string
	FlightNumber         string
	CarrierCode          string
	SegmentNumber        int
	StopoverIndicator    bool
	DepartureDateTime    time.Time
	ArrivalDateTime      time.Time
	BaggageAllowance     int
	StatusCode           string
	MaxPermittedMileage  int
	SegmentExtra         FareSegmentExtra
}

type FareQuote struct {
	TravelerCount                int
	FareNumber                   string
	Type                         string
	ValidatingCarrier            string
	PrivateAccountCode           string
	PrivateContractCode          string
	PrivateOfficeID              string
	ExcludeBasicEconomy          bool
	TaxExempt                    []string
	SegmentNumbers               []int
	FareBranding                 map[string]string
	Travelers                    []Traveler
	TourCode                     string
	CommissionPercent            float64
	CommissionDollar             float64
	TicketDesignator             string
	Endorsements                 string
	HemisphereCode               string
	JourneyCode                  string
	Payment                      Payment
	PenaltyApplies               bool
	Restrictions                 FareRestriction
	IsMileageFare                bool
	IsPrivateFare                bool
	HasCommissionPercentModifier bool
	HasCommissionDollarModifier  bool
	IsPhaseIVSabre               bool
	IsManualFare                 bool
	BaseFare                     Rate
	EquivalentFare               Rate
	TotalFare                    Rate
	ExchangeRate                 float64
	TaxList                      []FareTax
	InternationalSaleIndicator   string
	AgentID                      string
	Segments                     []FareSegment
	FareCalculation              string
	QuoteDate                    time.Time
	LastDateToTicket             time.Time
	IsTaxExempt                  bool
	InputMessage                 string
	AdditionalData               map[string]string
	FareText                     string
	Source                       Source
}

type Rate struct {
	Amount   float64
	Currency string
}

// Existing Pnr struct and other related structs...

type Pnr struct {
	Id                   string
	This                 AssociatedRecord
	AssociatedRecords    []AssociatedRecord
	BookingAgent         Agent
	LastQueueingAgent    Agent
	PhoneNumbers         []PhoneNumber
	EmailContacts        []EmailContact
	Travelers            []Traveler
	BillingAddress       Address
	DeliveryAddress      Address
	QueueHistory         []QueueHist
	AgencyAddress        Address
	Payments             []Payment
	Branch               string
	CustomerNumber       string
	ParentCustomerNumber string
	ETickets             []ETicket
	IsTicketed           bool
	McoLines             []McoLine
	SsrOsis              []Ssr
	TicketRequests       []TicketRequest
	BeginDate            time.Time
	EndDate              time.Time
	FutureQueueing       []FutureQueue
	AncillaryServices    []AncillaryServices
	AirSegments          []AirSegment
	Remark               Remark
	Remarks              []Remark
	AccountingRemark     Remark
	AccountingRemarks    []Remark
	PricingCart          []CartItem
	Invoices             []Invoice
	Communications       []Communication
}

// Other structs referenced in Pnr

type AssociatedRecord struct {
	RecordLocator      string
	OwningOffice       Office
	CreationOffice     Office
	CorporateProfileId string
	TravelerProfileId  string
	Source             Source
}

type AgentType int

const (
	Online AgentType = iota
	Onsite
	FullService
)

// String returns the string representation of the AgentType
func (at AgentType) String() string {
	return [...]string{"Online", "Onsite", "FullService"}[at]
}

type Agent struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	AgentType AgentType
	Queue     Queue
	Source    Source
}

type PhoneType int

const (
	Mobile PhoneType = iota
	Home
	Business
	Agency
)

func (p PhoneType) String() string {
	return [...]string{"Mobile", "Home", "Business", "Agency"}[p]
}

type PhoneNumber struct {
	CityCode    string
	CountryCode string
	Number      string
	Type        PhoneType
	LongText    string
}

type EmailContact struct {
	LineNumber int
	Address    string
	Extra      string
}

type TravelerType int

const (
	ADT TravelerType = iota
	CHD
	INF
)

func (t TravelerType) String() string {
	return [...]string{"ADT", "CHD", "INF"}[t]
}

type Traveler struct {
	Name             Name
	DateOfBirth      time.Time
	Type             TravelerType
	StatementIdField string
	IsInfantOnLap    bool
	NameNumber       NameNumber
	PassportNumber   string
	Nationality      string
	PassportTsa      string
	ProfileLocator   string
	Seats            []Seat
	FrequentFlyers   []FrequentFlyer
	SpecialMeals     []SpecialMeal
}

type Address struct {
	StreetAddress1 string
	StreetAddress2 string
	City           string
	StateProvince  StateProvince
	Country        Country
	PostCode       string
}

type QueueHist struct {
	Agent    Agent
	DateTime time.Time
	Queues   []Queue
	Source   Source
}

type PaymentType int

const (
	CCCard PaymentType = iota
	Cash
	Check
	InvoicePayment
)

type CreditCard struct {
	Code        string
	Number      string
	ExpiryMonth int
	ExpiryYear  int
}

// String returns the string representation of the PaymentType
func (pt PaymentType) String() string {
	return [...]string{"CreditCard", "Cash", "Check", "Invoice"}[pt]
}

type Payment struct {
	LineNumber    int
	TripTypeCodes []string
	UseTypeCodes  []string
	Type          PaymentType
	CreditCard    CreditCard
}

type ETicket struct {
	TicketNumber string
	Traveler     Traveler
}

type McoLine struct {
	Amount         float64
	Carrier        string
	IssuingCarrier string
	OfficeId       string
	Traveler       Traveler
	TicketNumber   string
	Description    string
	SegmentNumber  int
}

type Ssr struct {
	LineNumber        int
	SegmentNumbers    []int
	TravelerNumber    int
	Type              string
	AssociatedAirline string
	LongText          string
	Source            Source
}

type TicketRequest struct {
	LineNumber     int
	IsTicketed     bool
	BranchOffice   string
	RequestDate    time.Time
	FreeText       string
	AgentSine      string
	SegmentNumbers []int
}

type FQType int

const (
	Warning FQType = iota
	Cancel
)

// String returns the string representation of the FQType
func (fq FQType) String() string {
	return [...]string{"Warning", "Cancel"}[fq]
}

type Tax struct {
	Code     string
	Amount   float64
	Currency string
}

type FutureQueue struct {
	Queue          Queue
	Date           time.Time
	SegmentNumbers map[string]string
	FreeText       string
	Type           FQType
}

type AncillaryServices struct {
	Amount         float64
	DocumentNumber string
	SegmentNumber  int
	Flight         string
	Description    string
	Traveler       Traveler
	TotalAmount    float64
	Taxes          []Tax
	Source         Source
}

type Cabin int

const (
	Economy Cabin = iota
	PremiumEconomy
	BusinessClass
	FirstClass
)

// String returns the string representation of the Cabin
func (c Cabin) String() string {
	return [...]string{"Economy", "PremiumEconomy", "Business", "First"}[c]
}

type AirSegment struct {
	IsWaitlisted          bool
	IsCodeshare           bool
	EquipmentCode         string
	DayOfWeek             int
	CheckInDateTime       time.Time
	Carrier               string
	FlightNumber          string
	MealOnFlight          string
	HasSpecialMeal        bool
	SpecialMeal           string
	NumberInParty         int
	DepartureDateTime     time.Time
	ArrivalDateTime       time.Time
	DepartureAirportCode  string
	ArrivalAirportCode    string
	UtcArrivalDateTime    time.Time
	UtcDepartureDateTime  time.Time
	CalculatedDuration    string
	DepartureTerminal     string
	FlightDistance        Distance
	Stops                 []string
	ClassOfService        string
	ActiveRecord          AssociatedRecord
	NativeRecordLocator   string
	Status                string
	CodeShareFlightNumber string
	CodeShareCarrier      string
	ArrivalTerminal       string
	Flifo                 string
	CodeShareComment1     string
	CodeShareComment2     string
	IsChangeOfPlane       bool
	IsFlown               bool
	IsPassive             bool
	IsTicketless          bool
	Cabin                 Cabin
	DepartureCityCode     string
	ArrivalCityCode       string
	VendorName            string
	IsScheduleChange      bool
	IsCancelled           bool
	SegmentNumber         int
	AdditionalData        map[string]string
	Source                Source
}

type RemarkType int

const (
	GeneralRemark RemarkType = iota + 1
	InvoiceRemark
	ItineraryRemark
)

func (r RemarkType) String() string {
	return [...]string{"General", "Invoice", "Itinerary"}[r]
}

type Remark struct {
	IsCached        bool
	TravelerNumbers []int
	SegmentNumbers  []int
	Type            RemarkType
	Category        string
	Contents        string
	LineNumber      int
	CreatedDate     time.Time
	Source          Source
}

type CartDocType int

const (
	Fare CartDocType = iota
	Exchange
	Fee
	Passive
)

func (c CartDocType) String() string {
	return [...]string{"Fare", "Exchange", "Fee", "Passive"}[c]
}

type CartServiceType int

const (
	Air CartServiceType = iota
	ServiceFee
	Rail
	Tour
	Car
	Hotel
	GroundService
)

func (c CartServiceType) String() string {
	return [...]string{"Air", "Fee", "Rail", "Tour", "Car", "Hotel", "GroundService"}[c]
}

type CartItem struct {
	DocumentType     CartDocType
	ServiceType      CartServiceType
	Traveler         Traveler
	FareReference    string
	ExchangeNumber   string
	NonArcAccounting string
	FareQuote        FareQuote
	Payment          Payment
	State            string
	Name             string
	Tax              float64
	Document         string
	Type             string
	Base             Rate
	Total            float64
	Exchange         ExchangeItem
}

type Invoice struct {
	Number string
	Total  float64
	Date   time.Time
	Item   InvoiceItem
	Source Source
}

type CommType int

const (
	InvoiceComm CommType = iota
	Itinerary
	SelfService
	Notification
)

func (c CommType) String() string {
	return [...]string{"Invoice", "Itinerary", "SelfService", "Notification"}[c]
}

type Communication struct {
	Type    CommType
	Subject string
	CommId  string
	To      string
	Cc      string
	Bcc     string
	Status  string
}
