package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/imroc/req"
	"io"
	"strings"
)

type DashboardWidget struct {
	GUID                        *string   `json:"guid"`
	Title                       *string   `json:"title,omitempty"`
	Type                        *string   `json:"type"`
	Height                      *int      `json:"height"`
	Width                       *int      `json:"width"`
	X                           *int      `json:"x"`
	Y                           *int      `json:"y"`
	Label                       *string   `json:"label,omitempty"`
	AdqlQueries                 []*string `json:"adqlQueries,omitempty"`
	AnalyticsType               *string   `json:"analyticsType,omitempty"`
	SearchMode                  *string   `json:"searchMode,omitempty"`
	Description                 *string   `json:"description,omitempty"`
	DrillDownUrl                *string   `json:"drillDownUrl,omitempty"`
	UseMetricBrowserAsDrillDown *bool     `json:"useMetricBrowserAsDrillDown,omitempty"`
	DrillDownActionType         *string   `json:"drillDownActionType,omitempty"`
	BackgroundColor             *int      `json:"backgroundColor,omitempty"`
	Color                       *int      `json:"color,omitempty"`
	FontSize                    *int      `json:"fontSize,omitempty"`
	UseAutomaticFontSize        *bool     `json:"useAutomaticFontSize,omitempty"`
	BorderEnabled               *bool     `json:"borderEnabled,omitempty"`
	BorderThickness             *int      `json:"borderThickness,omitempty"`
	BorderColor                 *int      `json:"borderColor,omitempty"`
	BackgroundAlpha             float32   `json:"backgroundAlpha,omitempty"`
	ShowValues                  *bool     `json:"showValues,omitempty"`
	FormatNumber                *bool     `json:"formatNumber,omitempty"`
	NumDecimals                 *int      `json:"numDecimals,omitempty"`
	RemoveZeros                 *bool     `json:"removeZeros,omitempty"`
	BackgroundColors            []int     `json:"backgroundColors,omitempty"`
	CompactMode                 *bool     `json:"compactMode,omitempty"`
	ShowTimeRange               *bool     `json:"showTimeRange,omitempty"`
	RenderIn3D                  *bool     `json:"renderIn3D,omitempty"`
	ShowLegend                  *bool     `json:"showLegend,omitempty"`
	LegendPosition              *string   `json:"legendPosition,omitempty"`
	LegendColumnCount           *int      `json:"legendColumnCount,omitempty"`
	StartTime                   *bool     `json:"startTime,omitempty"`
	EndTime                     *int      `json:"endTime,omitempty"`
	CustomTimeRange             *int      `json:"customTimeRange,omitempty"`
	MinutesBeforeAnchorTime     *int      `json:"minutesBeforeAnchorTime,omitempty"`
	MinHeight                   *int      `json:"minHeight,omitempty"`
	MinWidth                    *int      `json:"minWidth,omitempty"`
	IsGlobal                    *bool     `json:"isGlobal,omitempty"`
	Resolution                  *string   `json:"resolution,omitempty"`

	IsShowLogYAxis           *bool    `json:"isShowLogYAxis,omitempty"`
	IsStackingEnabled        *bool    `json:"isStackingEnabled,omitempty"`
	LegendsLayout            *string  `json:"legendsLayout,omitempty"`
	MaxAllowedYAxisFields    *int     `json:"maxAllowedYAxisFields,omitempty"`
	MaxAllowedXAxisFields    *int     `json:"maxAllowedXAxisFields,omitempty"`
	ShowMinExtremes          *bool    `json:"showMinExtremes,omitempty"`
	ShowMaxExtremes          *bool    `json:"showMaxExtremes,omitempty"`
	DisplayPercentileMarkers *bool    `json:"displayPercentileMarkers,omitempty"`
	Unit                     *int     `json:"unit,omitempty"`
	IsRawQuery               *bool    `json:"isRawQuery,omitempty"`
	Align                    *string  `json:"align,omitempty"`
	ShowInverse              *bool    `json:"showInverse,omitempty"`
	ShowHealth               *bool    `json:"showHealth,omitempty"`
	IsIncreaseGood           *bool    `json:"isIncreaseGood,omitempty"`
	ShowUnivariateLabel      *bool    `json:"showUnivariateLabel,omitempty"`
	Properties               []string `json:"properties,omitempty"`
	VerticalAxisLabel        *string  `json:"verticalAxisLabel,omitempty"`
	HideHorizontalAxis       *bool    `json:"hideHorizontalAxis,omitempty"`
	HorizontalAxisLabel      *string  `json:"horizontalAxisLabel,omitempty"`
	AxisType                 *string  `json:"axisType,omitempty"`
	MultipleYAxis            *bool    `json:"multipleYAxis,omitempty"`
	CustomVerticalAxisMin    *bool    `json:"customVerticalAxisMin,omitempty"`
	CustomVerticalAxisMax    *bool    `json:"customVerticalAxisMax,omitempty"`
	ShowEvents               *bool    `json:"showEvents,omitempty"`
	EventFilter              *bool    `json:"eventFilter,omitempty"`
	InterpolateDataGaps      *bool    `json:"interpolateDataGaps,omitempty"`
	ShowAllTooltips          *bool    `json:"showAllTooltips,omitempty"`
	StaticThresholds         *string  `json:"staticThresholds,omitempty"`

	Text      *string `json:"text,omitempty"`
	TextAlign *string `json:"textAlign,omitempty"`
	Margin    *int    `json:"margin,omitempty"`
	// TODO
	//WidgetsMetricMatchCriterias null.String `json:"widgetsMetricMatchCriterias"`
	//MissingEntities     null.String `json:"missingEntities"`
}

type Dashboard struct {
	ID                      int               `json:"id"`
	Name                    string            `json:"name"`
	Width                   int               `json:"width"`
	Height                  int               `json:"height"`
	CanvasType              string            `json:"canvasType"`
	TemplateEntityType      string            `json:"templateEntityType"`
	RefreshInterval         int               `json:"refreshInterval"`
	BackgroundColor         int               `json:"backgroundColor"`
	WarRoom                 bool              `json:"warRoom"`
	Template                bool              `json:"template"`
	Widgets                 []DashboardWidget `json:"widgets"`
	Version                 int               `json:"version"`
	MinutesBeforeAnchorTime int               `json:"minutesBeforeAnchorTime"`
	StartTime               int               `json:"startTime"`
	EndTime                 int               `json:"endTime"`
}

type ImportExportResponse struct {
	Success   bool      `json:"success"`
	Dashboard Dashboard `json:"dashboard"`
}

func (c *AppDClient) createDashboardBaseUrl() string {
	return fmt.Sprintf("%s/controller/restui/dashboards", c.BaseUrl)
}

func (c *AppDClient) createDashboardUrl() string {
	return fmt.Sprintf("%s/createDashboard", c.createDashboardBaseUrl())
}

func (c *AppDClient) updateDashboardUrl() string {
	return fmt.Sprintf("%s/updateDashboard", c.createDashboardBaseUrl())
}

func (c *AppDClient) deleteDashboard() string {
	return fmt.Sprintf("%s/deleteDashboard", c.createDashboardBaseUrl())
}

func (c *AppDClient) deleteDashboards() string {
	return fmt.Sprintf("%s/deleteDashboards", c.createDashboardBaseUrl())
}

func (c *AppDClient) getDashboardUrl(dashboardId int) string {
	return fmt.Sprintf("%s/dashboardIfUpdated/%d/-1", c.createDashboardBaseUrl(), dashboardId)
}

func (c *AppDClient) importDashboardUrl() string {
	return fmt.Sprintf("%s/controller/CustomDashboardImportExportServlet", c.BaseUrl)
}

func (c *AppDClient) exportDashboardUrl(dashboardId int) string {
	return fmt.Sprintf("%s/controller/CustomDashboardImportExportServlet?dashboardId=%d", c.BaseUrl, dashboardId)
}

func (c *AppDClient) ImportDashboard(templateJson string) (*Dashboard, error) {
	file := io.NopCloser(strings.NewReader(templateJson))
	resp, err := req.Post(c.importDashboardUrl(), c.createAuthHeader(), req.FileUpload{
		File:      file,
		FieldName: "fileUpload",    // FieldName is form field name
		FileName:  "template.json", //Filename is the name of the file that you wish to upload. We use this to guess the mimetype as well as pass it onto the server
	})

	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error during dashboard import: %d, %s", resp.Response().StatusCode, respString))
	}

	appdResponse := ImportExportResponse{}
	err = resp.ToJSON(&appdResponse)

	return &appdResponse.Dashboard, err
}

func (c *AppDClient) CreateDashboard(dashboard Dashboard) (*Dashboard, error) {
	dashboardJson, err := json.Marshal(dashboard)
	if err != nil {
		return nil, err
	}
	resp, err := req.Post(c.createDashboardUrl(), c.createAuthHeader(), dashboardJson)
	if resp.Response().StatusCode != 200 {

		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error creating Dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	updated := Dashboard{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}
	return &updated, err
}

func (c *AppDClient) UpdateDashboard(dashboard Dashboard) (*Dashboard, error) {
	dashboardJson, err := json.Marshal(dashboard)
	if err != nil {
		return nil, err
	}
	resp, err := req.Post(c.updateDashboardUrl(), c.createAuthHeader(), dashboardJson)
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error updating Dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	updated := Dashboard{}
	err = resp.ToJSON(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, err
}

func (c *AppDClient) DeleteDashboard(dashboardId int) error {
	resp, err := req.Post(c.deleteDashboard(), c.createAuthHeader(), req.BodyJSON(dashboardId))
	if err != nil {
		return err
	}

	if resp.Response().StatusCode != 204 {
		respString, _ := resp.ToString()
		return errors.New(fmt.Sprintf("Error deleting Dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	return nil
}

func (c *AppDClient) GetDashboard(dashboardId int) (*Dashboard, error) {
	resp, err := req.Get(c.getDashboardUrl(dashboardId), c.createAuthHeader())
	if err != nil {
		return nil, err
	}
	if resp.Response().StatusCode != 200 {
		respString, _ := resp.ToString()
		return nil, errors.New(fmt.Sprintf("Error getting dashboard: %d, %s", resp.Response().StatusCode, respString))
	}
	dashboard := Dashboard{}
	err = resp.ToJSON(&dashboard)

	if err != nil {
		return nil, err
	}
	dashboard.ID = dashboardId
	return &dashboard, nil
}
