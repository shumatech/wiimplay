package upnp

type innerXml struct {
    InnerXml    string `xml:",innerxml"`
}

type DidlLiteXml struct {
    Items []MetadataXml `xml:"item"`
}

type MetadataXml struct {
    Id           string `xml:"id"`
    SubId        string `xml:"subid"`
    Description  string `xml:"description"`
    SkipLimit    int    `xml:"sliplimit"`
    Like         int    `xml:"like"`
    Res          string `xml:"res"`
    Title        string `xml:"title"`
    Album        string `xml:"album"`
    Artist       string `xml:"artist"`
    Creator      string `xml:"creator"`
    ThumbRating  string `xml:"thumbRating"`
    RatingUri    string `xml:"ratingURI"`
    AlbumArt     string `xml:"albumArtURI"`
    Rate         int    `xml:"rate_hz"`
    Format       string `xml:"format_s"`
    Quality      string `xml:"actualQuality"`
    BitRate      int    `xml:"bitrate"`
}

type TotalPlayQueueXml struct {
    TotalQueue      int      `xml:"TotalQueue"`
    CurrentPlayList string   `xml:"CurrentPlayList>Name"`
    PlayListInfo    innerXml `xml:"PlayListInfo"`
}

type TotalPlayListXml struct {
    Name     string               `xml:"Name"`
    ListInfo TotalPlayListInfoXml `xml:"ListInfo"`
}

type TotalPlayListInfoXml struct {
    Source         string `xml:"Source"`
    SearchUrl      string `xml:"SearchUrl"`
    AutoGenerate   int    `xml:"AutoGenerate"`
    StationLimit   int    `xml:"StationLimit"`
    MarkSearch     int    `xml:"MarkSearch"`
    Quality        int    `xml:"Quality"`
    UpdateTime     int    `xml:"UpdateTime"`
    LastPlayIndex  int    `xml:"LastPlayIndex"`
    AlarmPlayIndex int    `xml:"AlarmPlayIndex"`
    RealIndex      int    `xml:"RealIndex"`
    TrackNumber    int    `xml:"TrackNumber"`
    SwitchPageMode int    `xml:"SwitchPageMode"`
    PressType      int    `xml:"PressType"`
    Volume         int    `xml:"Volume"`
}

type PlayListXml struct {
    ListName string          `xml:"ListName"`
    ListInfo PlayListInfoXml `xml:"ListInfo"`
    Tracks   innerXml        `xml:"Tracks"`
}

type PlayListInfoXml struct {
    SourceName     string `xml:"SourceName"`
    MarkSearch     int    `xml:"MarkSearch"`
    TrackNumber    int    `xml:"TrackNumber"`
    Quality        int    `xml:"Quality"`
    UpdateTime     int    `xml:"UpdateTime"`
    LastPlayIndex  int    `xml:"LastPlayIndex"`
    AlarmPlayIndex int    `xml:"AlarmPlayIndex"`
    RealIndex      int    `xml:"RealIndex"`
    Type           int    `xml:"Type"`
    SwitchPageMode int    `xml:"SwitchPageMode"`
    CurrentPage    int    `xml:"CurrentPage"`
    TotalPages     int    `xml:"TotalPages"`
    Searching      int    `xml:"searching"`
    PressType      int    `xml:"PressType"`
    Volume         int    `xml:"Volume"`
    FadeEnable     int    `xml:"FadeEnable"`
    FadeInMS       int    `xml:"FadeInMS"`
    FadeOutMS      int    `xml:"FadeOutMS"`
}

type PlayListTrackXml struct {
    URL           string   `xml:"URL"`
    Metadata      string   `xml:"Metadata"`
    Id            string   `xml:"Id"`
    Source        string   `xml:"Source"`
    RefreshUrl    string   `xml:"RefreshUrl"`
    PlayEventUrl  string   `xml:"PlayEventUrl"`
    Expires       string   `xml:"Expires"`
    Key           string   `xml:"Key"`
    ChapterNumber int      `xml:"ChapterNumber"`
    Chapters      innerXml `xml:"Chapters"`
}