package main

type Event struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	City            string `json:"city"`
	Venue           string `json:"venue"`
	PublicationDate string `json:"publicationDate"`
	EventDate       string `json:"eventDate"`
}

type Attendee struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Price struct {
	Currency float64 `json:"currency"`
	Amount   string  `json:"amount"`
}

type Item struct {
	ID     int    `json:"id"`
	NameES string `json:"nameES"`
	NameEN string `json:"nameEN"`
	Price  Price  `json:"price"`
	Stock  int    `json:"stock"`
}

type Band struct {
	BandName string `json:"bandName"`
}

type Payment struct {
	ID         int    `json:"id"`
	ProviderID int    `json:"providerId"`
	AttendeeID int    `json:"attendeeId"`
	Concept    string `json:"concept"`
	DueTo      string `json:"dueTo"`
	URLReceipt string `json:"urlReceipt"`
	Price      Price  `json:"price"`
}

type Provider struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ItemsStadistics struct {
	TicketsSold             int    `json:"ticketsSold"`
	MerchandisingItemsSold  int    `json:"merchandisingItemsSold"`
	BestSellerMerchandising string `json:"bestSellerMerchandising"`
}

type ProvidersStadistics struct {
	ProviderID      int     `json:"providerId"`
	ProviderName    int     `json:"providerName"`
	TotalPaidAmount float64 `json:"totalPaidAmount"`
}

type AttendeesStadistics struct {
	NumberOfAttendees int     `json:"numberOfAttendees"`
	GrossRevenue      float64 `json:"grossRevenue"`
	Currency          int     `json:"currency"`
}

type EventStadistics struct {
	NumberOfAttendees int     `json:"numberOfAttendees"`
	GrossRevenue      float64 `json:"grossRevenue"`
	Costs             float64 `json:"costs"`
	Currency          string  `json:"currency"`
}
