package leankit

// BoardResponse response from api /board leankit
type BoardResponse struct {
	ReplyData []struct {
		IconConfiguration struct {
			IconPathTemplate  string `json:"IconPathTemplate"`
			AllowedIconImages []struct {
				Size         string `json:"Size"`
				OutputFormat string `json:"OutputFormat"`
			} `json:"AllowedIconImages"`
		} `json:"IconConfiguration"`
		ParentCards []struct {
			HasPermission   bool        `json:"HasPermission"`
			IsArchived      bool        `json:"IsArchived"`
			BoardID         int         `json:"BoardId"`
			IsCardIDEnabled bool        `json:"IsCardIdEnabled"`
			IsHeaderEnabled bool        `json:"IsHeaderEnabled"`
			IsPrefixEnabled bool        `json:"IsPrefixEnabled"`
			Prefix          interface{} `json:"Prefix"`
			SystemType      string      `json:"SystemType"`
			LaneID          int         `json:"LaneId"`
			Title           string      `json:"Title"`
			Description     string      `json:"Description"`
			Type            struct {
				ID int `json:"Id"`
			} `json:"Type"`
			TypeID                 int           `json:"TypeId"`
			Priority               int           `json:"Priority"`
			PriorityText           string        `json:"PriorityText"`
			TypeName               string        `json:"TypeName"`
			TypeIconPath           string        `json:"TypeIconPath"`
			TypeColorHex           string        `json:"TypeColorHex"`
			Size                   int           `json:"Size"`
			Active                 bool          `json:"Active"`
			Color                  string        `json:"Color"`
			Version                int           `json:"Version"`
			AssignedUsers          []interface{} `json:"AssignedUsers"`
			CardDrillThroughBoards []struct {
				ID int `json:"Id"`
			} `json:"CardDrillThroughBoards"`
			CardContexts                  interface{}   `json:"CardContexts"`
			IsBlocked                     bool          `json:"IsBlocked"`
			BlockReason                   string        `json:"BlockReason"`
			BlockStateChangeDate          string        `json:"BlockStateChangeDate"`
			Index                         int           `json:"Index"`
			DueDate                       string        `json:"DueDate"`
			StartDate                     string        `json:"StartDate"`
			ExternalSystemName            string        `json:"ExternalSystemName"`
			ExternalSystemURL             string        `json:"ExternalSystemUrl"`
			ExternalCardID                string        `json:"ExternalCardID"`
			Tags                          string        `json:"Tags"`
			CountOfOldCards               int           `json:"CountOfOldCards"`
			LastMove                      string        `json:"LastMove"`
			LastActivity                  string        `json:"LastActivity"`
			DateArchived                  string        `json:"DateArchived"`
			LastComment                   string        `json:"LastComment"`
			CommentsCount                 int           `json:"CommentsCount"`
			LastAttachment                string        `json:"LastAttachment"`
			AttachmentsCount              int           `json:"AttachmentsCount"`
			AssignedUserName              string        `json:"AssignedUserName"`
			AssignedUserID                int           `json:"AssignedUserId"`
			AssignedUserIds               []interface{} `json:"AssignedUserIds"`
			GravatarLink                  string        `json:"GravatarLink"`
			SmallGravatarLink             string        `json:"SmallGravatarLink"`
			CardDrillThroughBoardIds      []int         `json:"CardDrillThroughBoardIds"`
			ID                            int           `json:"Id"`
			DrillThroughBoardID           int           `json:"DrillThroughBoardId"`
			HasDrillThroughBoard          bool          `json:"HasDrillThroughBoard"`
			HasMultipleDrillThroughBoards bool          `json:"HasMultipleDrillThroughBoards"`
			DrillThroughStatistics        struct {
				ParentCard struct {
					ID                int         `json:"Id"`
					IsBlocked         bool        `json:"IsBlocked"`
					Size              int         `json:"Size"`
					PlannedStartDate  interface{} `json:"PlannedStartDate"`
					PlannedFinishDate interface{} `json:"PlannedFinishDate"`
					ActualStartDate   string      `json:"ActualStartDate"`
					ActualFinishDate  interface{} `json:"ActualFinishDate"`
				} `json:"ParentCard"`
				TotalNumberOfCards                                              int         `json:"TotalNumberOfCards"`
				NumberOfCardsNotStarted                                         int         `json:"NumberOfCardsNotStarted"`
				NumberOfCardsStarted                                            int         `json:"NumberOfCardsStarted"`
				NumberOfCardsCompleted                                          int         `json:"NumberOfCardsCompleted"`
				BlockedCards                                                    int         `json:"BlockedCards"`
				NumberOfCardsPastDue                                            int         `json:"NumberOfCardsPastDue"`
				ProjectedStartDate                                              interface{} `json:"ProjectedStartDate"`
				ProjectedFinishDate                                             interface{} `json:"ProjectedFinishDate"`
				ActualStartDate                                                 interface{} `json:"ActualStartDate"`
				ActualFinishDate                                                interface{} `json:"ActualFinishDate"`
				IsAnyCardProjectedFinishDatePastParentCardProjectedFinishedDate bool        `json:"IsAnyCardProjectedFinishDatePastParentCardProjectedFinishedDate"`
				TotalSizeOfCards                                                int         `json:"TotalSizeOfCards"`
				TotalSizeOfCardsNotStarted                                      int         `json:"TotalSizeOfCardsNotStarted"`
				TotalSizeOfCardsStarted                                         int         `json:"TotalSizeOfCardsStarted"`
				TotalSizeOfCardsCompleted                                       int         `json:"TotalSizeOfCardsCompleted"`
				TotalProgressPercentage                                         float64     `json:"TotalProgressPercentage"`
			} `json:"DrillThroughStatistics"`
			DrillThroughCompletionPercent    interface{}   `json:"DrillThroughCompletionPercent"`
			DrillThroughProgressTotal        int           `json:"DrillThroughProgressTotal"`
			DrillThroughProgressComplete     int           `json:"DrillThroughProgressComplete"`
			DrillThroughProgressSizeComplete int           `json:"DrillThroughProgressSizeComplete"`
			DrillThroughProgressSizeTotal    int           `json:"DrillThroughProgressSizeTotal"`
			ClassOfServiceID                 int           `json:"ClassOfServiceId"`
			ClassOfServiceTitle              string        `json:"ClassOfServiceTitle"`
			ClassOfServiceIconPath           string        `json:"ClassOfServiceIconPath"`
			ClassOfServiceColorHex           string        `json:"ClassOfServiceColorHex"`
			ClassOfServiceCustomIconName     interface{}   `json:"ClassOfServiceCustomIconName"`
			ClassOfServiceCustomIconColor    interface{}   `json:"ClassOfServiceCustomIconColor"`
			CardTypeIconColor                interface{}   `json:"CardTypeIconColor"`
			CardTypeIconName                 interface{}   `json:"CardTypeIconName"`
			CurrentTaskBoardID               interface{}   `json:"CurrentTaskBoardId"`
			TaskBoardCompletionPercent       interface{}   `json:"TaskBoardCompletionPercent"`
			TaskBoardCompletedCardCount      int           `json:"TaskBoardCompletedCardCount"`
			TaskBoardCompletedCardSize       int           `json:"TaskBoardCompletedCardSize"`
			TaskBoardTotalCards              interface{}   `json:"TaskBoardTotalCards"`
			TaskBoardTotalSize               interface{}   `json:"TaskBoardTotalSize"`
			CurrentContext                   interface{}   `json:"CurrentContext"`
			ParentCardID                     interface{}   `json:"ParentCardId"`
			ParentCardIds                    []interface{} `json:"ParentCardIds"`
		} `json:"ParentCards"`
		ID                     int         `json:"Id"`
		Title                  string      `json:"Title"`
		Description            interface{} `json:"Description"`
		Version                int         `json:"Version"`
		Active                 bool        `json:"Active"`
		OrganizationID         int         `json:"OrganizationId"`
		OrganizationActivities []struct {
			ID   int    `json:"Id"`
			Name string `json:"Name"`
		} `json:"OrganizationActivities"`
		Lanes []struct {
			ID            int         `json:"Id"`
			Description   string      `json:"Description"`
			Index         int         `json:"Index"`
			Active        bool        `json:"Active"`
			Title         string      `json:"Title"`
			CardLimit     int         `json:"CardLimit"`
			ClassType     int         `json:"ClassType"`
			Type          int         `json:"Type"`
			ActivityID    interface{} `json:"ActivityId"`
			ActivityName  string      `json:"ActivityName"`
			CardContextID int         `json:"CardContextId"`
			Width         int         `json:"Width"`
			ParentLaneID  int         `json:"ParentLaneId"`
			Cards         []struct {
				SystemType  string `json:"SystemType"`
				LaneID      int    `json:"LaneId"`
				Title       string `json:"Title"`
				Description string `json:"Description"`
				Type        struct {
					ID int `json:"Id"`
				} `json:"Type"`
				TypeID        int    `json:"TypeId"`
				Priority      int    `json:"Priority"`
				PriorityText  string `json:"PriorityText"`
				TypeName      string `json:"TypeName"`
				TypeIconPath  string `json:"TypeIconPath"`
				TypeColorHex  string `json:"TypeColorHex"`
				Size          int    `json:"Size"`
				Active        bool   `json:"Active"`
				Color         string `json:"Color"`
				Version       int    `json:"Version"`
				AssignedUsers []struct {
					ID                int    `json:"Id"`
					GravatarLink      string `json:"GravatarLink"`
					SmallGravatarLink string `json:"SmallGravatarLink"`
					FullName          string `json:"FullName"`
					EmailAddress      string `json:"EmailAddress"`
					HasGravatar       bool   `json:"HasGravatar"`
					AssignedUserName  string `json:"AssignedUserName"`
					AssignedUserID    int    `json:"AssignedUserId"`
				} `json:"AssignedUsers"`
				CardDrillThroughBoards           []interface{} `json:"CardDrillThroughBoards"`
				CardContexts                     interface{}   `json:"CardContexts"`
				IsBlocked                        bool          `json:"IsBlocked"`
				BlockReason                      string        `json:"BlockReason"`
				BlockStateChangeDate             string        `json:"BlockStateChangeDate"`
				Index                            int           `json:"Index"`
				DueDate                          string        `json:"DueDate"`
				StartDate                        string        `json:"StartDate"`
				ExternalSystemName               string        `json:"ExternalSystemName"`
				ExternalSystemURL                string        `json:"ExternalSystemUrl"`
				ExternalCardID                   string        `json:"ExternalCardID"`
				Tags                             string        `json:"Tags"`
				CountOfOldCards                  int           `json:"CountOfOldCards"`
				LastMove                         string        `json:"LastMove"`
				LastActivity                     string        `json:"LastActivity"`
				DateArchived                     string        `json:"DateArchived"`
				LastComment                      string        `json:"LastComment"`
				CommentsCount                    int           `json:"CommentsCount"`
				LastAttachment                   string        `json:"LastAttachment"`
				AttachmentsCount                 int           `json:"AttachmentsCount"`
				AssignedUserName                 string        `json:"AssignedUserName"`
				AssignedUserID                   int           `json:"AssignedUserId"`
				AssignedUserIds                  []int         `json:"AssignedUserIds"`
				GravatarLink                     string        `json:"GravatarLink"`
				SmallGravatarLink                string        `json:"SmallGravatarLink"`
				CardDrillThroughBoardIds         []interface{} `json:"CardDrillThroughBoardIds"`
				ID                               int           `json:"Id"`
				DrillThroughBoardID              interface{}   `json:"DrillThroughBoardId"`
				HasDrillThroughBoard             bool          `json:"HasDrillThroughBoard"`
				HasMultipleDrillThroughBoards    bool          `json:"HasMultipleDrillThroughBoards"`
				DrillThroughStatistics           interface{}   `json:"DrillThroughStatistics"`
				DrillThroughCompletionPercent    float64       `json:"DrillThroughCompletionPercent"`
				DrillThroughProgressTotal        int           `json:"DrillThroughProgressTotal"`
				DrillThroughProgressComplete     int           `json:"DrillThroughProgressComplete"`
				DrillThroughProgressSizeComplete int           `json:"DrillThroughProgressSizeComplete"`
				DrillThroughProgressSizeTotal    int           `json:"DrillThroughProgressSizeTotal"`
				ClassOfServiceID                 int           `json:"ClassOfServiceId"`
				ClassOfServiceTitle              string        `json:"ClassOfServiceTitle"`
				ClassOfServiceIconPath           string        `json:"ClassOfServiceIconPath"`
				ClassOfServiceColorHex           string        `json:"ClassOfServiceColorHex"`
				ClassOfServiceCustomIconName     interface{}   `json:"ClassOfServiceCustomIconName"`
				ClassOfServiceCustomIconColor    interface{}   `json:"ClassOfServiceCustomIconColor"`
				CardTypeIconColor                string        `json:"CardTypeIconColor"`
				CardTypeIconName                 string        `json:"CardTypeIconName"`
				CurrentTaskBoardID               interface{}   `json:"CurrentTaskBoardId"`
				TaskBoardCompletionPercent       interface{}   `json:"TaskBoardCompletionPercent"`
				TaskBoardCompletedCardCount      int           `json:"TaskBoardCompletedCardCount"`
				TaskBoardCompletedCardSize       int           `json:"TaskBoardCompletedCardSize"`
				TaskBoardTotalCards              interface{}   `json:"TaskBoardTotalCards"`
				TaskBoardTotalSize               interface{}   `json:"TaskBoardTotalSize"`
				CurrentContext                   interface{}   `json:"CurrentContext"`
				ParentCardID                     int           `json:"ParentCardId"`
				ParentCardIds                    []int         `json:"ParentCardIds"`
			} `json:"Cards"`
			Orientation            int         `json:"Orientation"`
			TaskBoardID            interface{} `json:"TaskBoardId"`
			ChildLaneIds           []int       `json:"ChildLaneIds"`
			SiblingLaneIds         []int       `json:"SiblingLaneIds"`
			IsDrillthroughDoneLane bool        `json:"IsDrillthroughDoneLane"`
			IsDefaultDropLane      bool        `json:"IsDefaultDropLane"`
			LaneState              string      `json:"LaneState"`
		} `json:"Lanes"`
		BoardUsers []struct {
			ID             int         `json:"Id"`
			FullName       string      `json:"FullName"`
			UserName       string      `json:"UserName"`
			Role           int         `json:"Role"`
			WIP            int         `json:"WIP"`
			Enabled        bool        `json:"Enabled"`
			IsAccountOwner bool        `json:"IsAccountOwner"`
			IsDeleted      bool        `json:"IsDeleted"`
			GravatarFeed   string      `json:"GravatarFeed"`
			GravatarLink   string      `json:"GravatarLink"`
			EmailAddress   string      `json:"EmailAddress"`
			DateFormat     string      `json:"DateFormat"`
			Settings       interface{} `json:"Settings"`
			RoleName       string      `json:"RoleName"`
		} `json:"BoardUsers"`
		Backlog []struct {
			ID            int         `json:"Id"`
			Description   string      `json:"Description"`
			Index         int         `json:"Index"`
			Active        bool        `json:"Active"`
			Title         string      `json:"Title"`
			CardLimit     int         `json:"CardLimit"`
			ClassType     int         `json:"ClassType"`
			Type          int         `json:"Type"`
			ActivityID    interface{} `json:"ActivityId"`
			ActivityName  string      `json:"ActivityName"`
			CardContextID int         `json:"CardContextId"`
			Width         int         `json:"Width"`
			ParentLaneID  int         `json:"ParentLaneId"`
			Cards         []struct {
				SystemType  string `json:"SystemType"`
				LaneID      int    `json:"LaneId"`
				Title       string `json:"Title"`
				Description string `json:"Description"`
				Type        struct {
					ID int `json:"Id"`
				} `json:"Type"`
				TypeID        int    `json:"TypeId"`
				Priority      int    `json:"Priority"`
				PriorityText  string `json:"PriorityText"`
				TypeName      string `json:"TypeName"`
				TypeIconPath  string `json:"TypeIconPath"`
				TypeColorHex  string `json:"TypeColorHex"`
				Size          int    `json:"Size"`
				Active        bool   `json:"Active"`
				Color         string `json:"Color"`
				Version       int    `json:"Version"`
				AssignedUsers []struct {
					ID                int    `json:"Id"`
					GravatarLink      string `json:"GravatarLink"`
					SmallGravatarLink string `json:"SmallGravatarLink"`
					FullName          string `json:"FullName"`
					EmailAddress      string `json:"EmailAddress"`
					HasGravatar       bool   `json:"HasGravatar"`
					AssignedUserName  string `json:"AssignedUserName"`
					AssignedUserID    int    `json:"AssignedUserId"`
				} `json:"AssignedUsers"`
				CardDrillThroughBoards           []interface{} `json:"CardDrillThroughBoards"`
				CardContexts                     interface{}   `json:"CardContexts"`
				IsBlocked                        bool          `json:"IsBlocked"`
				BlockReason                      string        `json:"BlockReason"`
				BlockStateChangeDate             string        `json:"BlockStateChangeDate"`
				Index                            int           `json:"Index"`
				DueDate                          string        `json:"DueDate"`
				StartDate                        string        `json:"StartDate"`
				ExternalSystemName               string        `json:"ExternalSystemName"`
				ExternalSystemURL                string        `json:"ExternalSystemUrl"`
				ExternalCardID                   string        `json:"ExternalCardID"`
				Tags                             string        `json:"Tags"`
				CountOfOldCards                  int           `json:"CountOfOldCards"`
				LastMove                         string        `json:"LastMove"`
				LastActivity                     string        `json:"LastActivity"`
				DateArchived                     string        `json:"DateArchived"`
				LastComment                      string        `json:"LastComment"`
				CommentsCount                    int           `json:"CommentsCount"`
				LastAttachment                   string        `json:"LastAttachment"`
				AttachmentsCount                 int           `json:"AttachmentsCount"`
				AssignedUserName                 string        `json:"AssignedUserName"`
				AssignedUserID                   int           `json:"AssignedUserId"`
				AssignedUserIds                  []int         `json:"AssignedUserIds"`
				GravatarLink                     string        `json:"GravatarLink"`
				SmallGravatarLink                string        `json:"SmallGravatarLink"`
				CardDrillThroughBoardIds         []interface{} `json:"CardDrillThroughBoardIds"`
				ID                               int           `json:"Id"`
				DrillThroughBoardID              interface{}   `json:"DrillThroughBoardId"`
				HasDrillThroughBoard             bool          `json:"HasDrillThroughBoard"`
				HasMultipleDrillThroughBoards    bool          `json:"HasMultipleDrillThroughBoards"`
				DrillThroughStatistics           interface{}   `json:"DrillThroughStatistics"`
				DrillThroughCompletionPercent    float64       `json:"DrillThroughCompletionPercent"`
				DrillThroughProgressTotal        int           `json:"DrillThroughProgressTotal"`
				DrillThroughProgressComplete     int           `json:"DrillThroughProgressComplete"`
				DrillThroughProgressSizeComplete int           `json:"DrillThroughProgressSizeComplete"`
				DrillThroughProgressSizeTotal    int           `json:"DrillThroughProgressSizeTotal"`
				ClassOfServiceID                 int           `json:"ClassOfServiceId"`
				ClassOfServiceTitle              string        `json:"ClassOfServiceTitle"`
				ClassOfServiceIconPath           string        `json:"ClassOfServiceIconPath"`
				ClassOfServiceColorHex           string        `json:"ClassOfServiceColorHex"`
				ClassOfServiceCustomIconName     interface{}   `json:"ClassOfServiceCustomIconName"`
				ClassOfServiceCustomIconColor    interface{}   `json:"ClassOfServiceCustomIconColor"`
				CardTypeIconColor                string        `json:"CardTypeIconColor"`
				CardTypeIconName                 string        `json:"CardTypeIconName"`
				CurrentTaskBoardID               interface{}   `json:"CurrentTaskBoardId"`
				TaskBoardCompletionPercent       interface{}   `json:"TaskBoardCompletionPercent"`
				TaskBoardCompletedCardCount      int           `json:"TaskBoardCompletedCardCount"`
				TaskBoardCompletedCardSize       int           `json:"TaskBoardCompletedCardSize"`
				TaskBoardTotalCards              interface{}   `json:"TaskBoardTotalCards"`
				TaskBoardTotalSize               interface{}   `json:"TaskBoardTotalSize"`
				CurrentContext                   interface{}   `json:"CurrentContext"`
				ParentCardID                     int           `json:"ParentCardId"`
				ParentCardIds                    []int         `json:"ParentCardIds"`
			} `json:"Cards"`
			Orientation            int         `json:"Orientation"`
			TaskBoardID            interface{} `json:"TaskBoardId"`
			ChildLaneIds           []int       `json:"ChildLaneIds"`
			SiblingLaneIds         []int       `json:"SiblingLaneIds"`
			IsDrillthroughDoneLane bool        `json:"IsDrillthroughDoneLane"`
			IsDefaultDropLane      bool        `json:"IsDefaultDropLane"`
			LaneState              string      `json:"LaneState"`
		} `json:"Backlog"`
		Archive []struct {
			ID                     int           `json:"Id"`
			Description            string        `json:"Description"`
			Index                  int           `json:"Index"`
			Active                 bool          `json:"Active"`
			Title                  string        `json:"Title"`
			CardLimit              int           `json:"CardLimit"`
			ClassType              int           `json:"ClassType"`
			Type                   int           `json:"Type"`
			ActivityID             interface{}   `json:"ActivityId"`
			ActivityName           string        `json:"ActivityName"`
			CardContextID          int           `json:"CardContextId"`
			Width                  int           `json:"Width"`
			ParentLaneID           int           `json:"ParentLaneId"`
			Cards                  []interface{} `json:"Cards"`
			Orientation            int           `json:"Orientation"`
			TaskBoardID            interface{}   `json:"TaskBoardId"`
			ChildLaneIds           []int         `json:"ChildLaneIds"`
			SiblingLaneIds         []int         `json:"SiblingLaneIds"`
			IsDrillthroughDoneLane bool          `json:"IsDrillthroughDoneLane"`
			IsDefaultDropLane      bool          `json:"IsDefaultDropLane"`
			LaneState              string        `json:"LaneState"`
		} `json:"Archive"`
		ClassesOfService []struct {
			ID              int         `json:"Id"`
			Title           string      `json:"Title"`
			Policy          string      `json:"Policy"`
			IconPath        string      `json:"IconPath"`
			ColorHex        string      `json:"ColorHex"`
			CustomIconName  interface{} `json:"CustomIconName"`
			CustomIconColor interface{} `json:"CustomIconColor"`
			UseColor        bool        `json:"UseColor"`
		} `json:"ClassesOfService"`
		CurrentUserRole int `json:"CurrentUserRole"`
		CardTypes       []struct {
			ID                int         `json:"Id"`
			Name              string      `json:"Name"`
			ColorHex          string      `json:"ColorHex"`
			IsDefault         bool        `json:"IsDefault"`
			IconPath          interface{} `json:"IconPath"`
			IsCardType        bool        `json:"IsCardType"`
			IsTaskType        bool        `json:"IsTaskType"`
			IsDefaultTaskType bool        `json:"IsDefaultTaskType"`
			IconName          string      `json:"IconName"`
			IconColor         string      `json:"IconColor"`
		} `json:"CardTypes"`
		ClassOfServiceEnabled                  bool        `json:"ClassOfServiceEnabled"`
		CardColorField                         string      `json:"CardColorField"`
		IsCardIDEnabled                        bool        `json:"IsCardIdEnabled"`
		IsHeaderEnabled                        bool        `json:"IsHeaderEnabled"`
		IsPrefixEnabled                        bool        `json:"IsPrefixEnabled"`
		IsPrefixIncludedInHyperlink            bool        `json:"IsPrefixIncludedInHyperlink"`
		Prefix                                 string      `json:"Prefix"`
		IsHyperlinkEnabled                     bool        `json:"IsHyperlinkEnabled"`
		Format                                 string      `json:"Format"`
		AllowMultiUserAssignments              bool        `json:"AllowMultiUserAssignments"`
		ArchiveCardDays                        int         `json:"ArchiveCardDays"`
		EnableDrillThroughBoards               bool        `json:"EnableDrillThroughBoards"`
		EnableMultipleDrillThroughBoards       bool        `json:"EnableMultipleDrillThroughBoards"`
		EnableDrillThroughDoneLane             bool        `json:"EnableDrillThroughDoneLane"`
		AllowMoveCardsBetweenBoards            bool        `json:"AllowMoveCardsBetweenBoards"`
		SubscribeUsersToAssignedCardsByDefault bool        `json:"SubscribeUsersToAssignedCardsByDefault"`
		AllowAttachments                       bool        `json:"AllowAttachments"`
		AllowComments                          bool        `json:"AllowComments"`
		AllowHorizontalSplitInBoardEdit        bool        `json:"AllowHorizontalSplitInBoardEdit"`
		MaxFileSize                            int         `json:"MaxFileSize"`
		DisallowedFileExtensions               interface{} `json:"DisallowedFileExtensions"`
		EnableTaskBoards                       bool        `json:"EnableTaskBoards"`
		AllowSeparateCardAndTaskTypes          bool        `json:"AllowSeparateCardAndTaskTypes"`
		AllowAddCardTypes                      bool        `json:"AllowAddCardTypes"`
		BaseWipOnCardSize                      bool        `json:"BaseWipOnCardSize"`
		AllowUsersToDeleteCards                bool        `json:"AllowUsersToDeleteCards"`
		ExcludeCompletedAndArchiveViolations   bool        `json:"ExcludeCompletedAndArchiveViolations"`
		IsWelcome                              bool        `json:"IsWelcome"`
		IsShared                               bool        `json:"IsShared"`
		SharedBoardRole                        int         `json:"SharedBoardRole"`
		CustomBoardMoniker                     interface{} `json:"CustomBoardMoniker"`
		TopLevelLaneIds                        []int       `json:"TopLevelLaneIds"`
		BacklogTopLevelLaneID                  int         `json:"BacklogTopLevelLaneId"`
		ArchiveTopLevelLaneID                  int         `json:"ArchiveTopLevelLaneId"`
		DefaultDropLaneID                      int         `json:"DefaultDropLaneId"`
		HasCustomFields                        bool        `json:"HasCustomFields"`
		CardContexts                           []struct {
			ID        int    `json:"Id"`
			Name      string `json:"Name"`
			IsDefault bool   `json:"IsDefault"`
			IsMaster  bool   `json:"IsMaster"`
		} `json:"CardContexts"`
		BoardCreator                 interface{} `json:"BoardCreator"`
		AvailableTags                string      `json:"AvailableTags"`
		IsAutoIncrementCardIDEnabled bool        `json:"IsAutoIncrementCardIdEnabled"`
		IsExternalURLEnabled         bool        `json:"IsExternalUrlEnabled"`
		IsPermalinkEnabled           bool        `json:"IsPermalinkEnabled"`
		CustomIconFieldLabel         string      `json:"CustomIconFieldLabel"`
		EnableLanePolicies           bool        `json:"EnableLanePolicies"`
		EnableLaneSubscription       bool        `json:"EnableLaneSubscription"`
		EnableCardHistory            bool        `json:"EnableCardHistory"`
		EnableWipLimit               bool        `json:"EnableWipLimit"`
	} `json:"ReplyData"`
	ReplyCode int    `json:"ReplyCode"`
	ReplyText string `json:"ReplyText"`
}
