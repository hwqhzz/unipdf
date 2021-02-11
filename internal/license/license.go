//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package license ;import (_ef "bytes";_d "compress/gzip";_e "crypto";_db "crypto/aes";_fa "crypto/cipher";_ab "crypto/rand";_def "crypto/rsa";_ada "crypto/sha512";_cc "crypto/x509";_bc "encoding/base64";_fg "encoding/hex";_adf "encoding/json";_ad "encoding/pem";_de "errors";_fb "fmt";_ga "github.com/unidoc/unipdf/v3/common";_g "io";_ag "io/ioutil";_gf "net";_be "net/http";_ed "os";_a "path/filepath";_c "sort";_efg "strings";_b "sync";_dc "time";);type LicenseKey struct{LicenseId string `json:"license_id"`;CustomerId string `json:"customer_id"`;CustomerName string `json:"customer_name"`;Tier string `json:"tier"`;CreatedAt _dc .Time `json:"-"`;CreatedAtInt int64 `json:"created_at"`;ExpiresAt *_dc .Time `json:"-"`;ExpiresAtInt int64 `json:"expires_at"`;CreatedBy string `json:"created_by"`;CreatorName string `json:"creator_name"`;CreatorEmail string `json:"creator_email"`;UniPDF bool `json:"unipdf"`;UniOffice bool `json:"unioffice"`;UniHTML bool `json:"unihtml"`;Trial bool `json:"trial"`;_dd bool ;_eaa string ;};func _facf (_abg string )(LicenseKey ,error ){var _edf LicenseKey ;_ea ,_eb :=_ee (_abe ,_bb ,_abg );if _eb !=nil {return _edf ,_eb ;};_bbd ,_eb :=_cb (_eaag ,_ea );if _eb !=nil {return _edf ,_eb ;};_eb =_adf .Unmarshal (_bbd ,&_edf );if _eb !=nil {return _edf ,_eb ;};_edf .CreatedAt =_dc .Unix (_edf .CreatedAtInt ,0);if _edf .ExpiresAtInt > 0{_adb :=_dc .Unix (_edf .ExpiresAtInt ,0);_edf .ExpiresAt =&_adb ;};return _edf ,nil ;};func (_bee *meteredClient )getStatus ()(meteredStatusResp ,error ){var _dcc meteredStatusResp ;_bcf :=_bee ._ddd +"\u002fm\u0065t\u0065\u0072\u0065\u0064\u002f\u0073\u0074\u0061\u0074\u0075\u0073";var _ce meteredStatusForm ;_cf ,_bdb :=_adf .Marshal (_ce );if _bdb !=nil {return _dcc ,_bdb ;};_fbd ,_bdb :=_ffcc (_cf );if _bdb !=nil {return _dcc ,_bdb ;};_eef ,_bdb :=_be .NewRequest ("\u0050\u004f\u0053\u0054",_bcf ,_fbd );if _bdb !=nil {return _dcc ,_bdb ;};_eef .Header .Add ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e");_eef .Header .Add ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_eef .Header .Add ("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_eef .Header .Add ("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y",_bee ._aeb );_cbg ,_bdb :=_bee ._ggea .Do (_eef );if _bdb !=nil {return _dcc ,_bdb ;};defer _cbg .Body .Close ();if _cbg .StatusCode !=200{return _dcc ,_fb .Errorf ("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064",_cbg .StatusCode );};_da ,_bdb :=_fggb (_cbg );if _bdb !=nil {return _dcc ,_bdb ;};_bdb =_adf .Unmarshal (_da ,&_dcc );if _bdb !=nil {return _dcc ,_bdb ;};return _dcc ,nil ;};const (LicenseTierUnlicensed ="\u0075\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";LicenseTierCommunity ="\u0063o\u006d\u006d\u0075\u006e\u0069\u0074y";LicenseTierIndividual ="\u0069\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";LicenseTierBusiness ="\u0062\u0075\u0073\u0069\u006e\u0065\u0073\u0073";);func _cad (_eda string ,_adg string ,_afa bool )error {if _aea ==nil {return _de .New ("\u006e\u006f\u0020\u006c\u0069\u0063\u0065\u006e\u0073e\u0020\u006b\u0065\u0079");};if !_aea ._dd ||len (_aea ._eaa )==0{return nil ;};if len (_eda )==0&&!_afa {return _de .New ("\u0064\u006f\u0063\u004b\u0065\u0079\u0020\u006e\u006ft\u0020\u0073\u0065\u0074");};_geb .Lock ();defer _geb .Unlock ();if _ecf ==nil {_ecf =map[string ]struct{}{};};if _eac ==nil {_eac =map[string ]int {};};_efc :=0;if !_afa {_ ,_ccd :=_ecf [_eda ];if !_ccd {_ecf [_eda ]=struct{}{};_efc ++;};if _efc ==0&&!_afa {return nil ;};_eac [_adg ]++;};_gfc :=_dc .Now ();_bfg ,_fae :=_beg .loadState (_aea ._eaa );if _fae !=nil {_ga .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_fae );return _fae ;};if _bfg .Usage ==nil {_bfg .Usage =map[string ]int {};};for _bbfe ,_dcb :=range _eac {_bfg .Usage [_bbfe ]+=_dcb ;};_eac =nil ;const _cded =24*_dc .Hour ;const _dcga =3*24*_dc .Hour ;if len (_bfg .Instance )==0||_gfc .Sub (_bfg .LastReported )> _cded ||(_bfg .LimitDocs &&_bfg .RemainingDocs <=_bfg .Docs +int64 (_efc ))||_afa {_fbb ,_fgdb :=_ed .Hostname ();if _fgdb !=nil {return _fgdb ;};_bgg :=_bfg .Docs ;_edgf ,_cagc ,_fgdb :=_cgf ();if _fgdb !=nil {return _fgdb ;};_c .Strings (_cagc );_c .Strings (_edgf );_gdf ,_fgdb :=_adc ();if _fgdb !=nil {return _fgdb ;};_dceb :=false ;for _ ,_gfe :=range _cagc {if _gfe ==_gdf .String (){_dceb =true ;};};if !_dceb {_cagc =append (_cagc ,_gdf .String ());};_ggac :=_ffc ();_ggac ._aeb =_aea ._eaa ;_bgg +=int64 (_efc );_cbag :=meteredUsageCheckinForm {Instance :_bfg .Instance ,Next :_bfg .Next ,UsageNumber :int (_bgg ),NumFailed :_bfg .NumErrors ,Hostname :_fbb ,LocalIP :_efg .Join (_cagc ,"\u002c\u0020"),MacAddress :_efg .Join (_edgf ,"\u002c\u0020"),Package :"\u0075\u006e\u0069\u0070\u0064\u0066",PackageVersion :_ga .Version ,Usage :_bfg .Usage };_bfd :=int64 (0);_gged :=_bfg .NumErrors ;_eagc :=_gfc ;_dbec :=0;_ccdc :=_bfg .LimitDocs ;_abdc ,_fgdb :=_ggac .checkinUsage (_cbag );if _fgdb !=nil {if _gfc .Sub (_bfg .LastReported )> _dcga {return _de .New ("\u0074\u006f\u006f\u0020\u006c\u006f\u006e\u0067\u0020\u0073\u0069\u006e\u0063\u0065\u0020\u006c\u0061\u0073\u0074\u0020\u0073\u0075\u0063\u0063e\u0073\u0073\u0066\u0075\u006c \u0063\u0068e\u0063\u006b\u0069\u006e");};_bfd =_bgg ;_gged ++;_eagc =_bfg .LastReported ;}else {_ccdc =_abdc .LimitDocs ;_dbec =_abdc .RemainingDocs ;_gged =0;};if len (_abdc .Instance )==0{_abdc .Instance =_cbag .Instance ;};if len (_abdc .Next )==0{_abdc .Next =_cbag .Next ;};_fgdb =_beg .updateState (_ggac ._aeb ,_abdc .Instance ,_abdc .Next ,int (_bfd ),_ccdc ,_dbec ,int (_gged ),_eagc ,nil );if _fgdb !=nil {return _fgdb ;};if !_abdc .Success {return _fb .Errorf ("\u0065r\u0072\u006f\u0072\u003a\u0020\u0025s",_abdc .Message );};}else {_fae =_beg .updateState (_aea ._eaa ,_bfg .Instance ,_bfg .Next ,int (_bfg .Docs )+_efc ,_bfg .LimitDocs ,int (_bfg .RemainingDocs ),int (_bfg .NumErrors ),_bfg .LastReported ,_bfg .Usage );if _fae !=nil {return _fae ;};};return nil ;};var _geb =&_b .Mutex {};func (_fbc *LicenseKey )Validate ()error {if len (_fbc .LicenseId )< 10{return _fb .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020L\u0069\u0063\u0065n\u0073e\u0020\u0049\u0064");};if len (_fbc .CustomerId )< 10{return _fb .Errorf ("\u0069\u006e\u0076\u0061l\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020C\u0075\u0073\u0074\u006f\u006d\u0065\u0072 \u0049\u0064");};if len (_fbc .CustomerName )< 1{return _fb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043u\u0073\u0074\u006f\u006d\u0065\u0072\u0020\u004e\u0061\u006d\u0065");};if _eba .After (_fbc .CreatedAt ){return _fb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064 \u0041\u0074\u0020\u0069\u0073 \u0069\u006ev\u0061\u006c\u0069\u0064");};if _fbc .ExpiresAt ==nil {_dee :=_fbc .CreatedAt .AddDate (1,0,0);if _gga .After (_dee ){_dee =_gga ;};_fbc .ExpiresAt =&_dee ;};if _fbc .CreatedAt .After (*_fbc .ExpiresAt ){return _fb .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0041\u0074 \u0063a\u006e\u006e\u006f\u0074 \u0062\u0065 \u0047\u0072\u0065\u0061\u0074\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0045\u0078\u0070\u0069\u0072\u0065\u0073\u0020\u0041\u0074");};if _fbc .isExpired (){return _fb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020l\u0069\u0063\u0065ns\u0065\u003a\u0020\u0054\u0068\u0065 \u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u0068\u0061\u0073\u0020\u0061\u006c\u0072e\u0061\u0064\u0079\u0020\u0065\u0078\u0070\u0069r\u0065\u0064");};if len (_fbc .CreatorName )< 1{return _fb .Errorf ("\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u003a\u0020\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u0020na\u006d\u0065");};if len (_fbc .CreatorEmail )< 1{return _fb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069c\u0065\u006e\u0073\u0065\u003a\u0020\u0043r\u0065\u0061\u0074\u006f\u0072\u0020\u0065\u006d\u0061\u0069\u006c");};if _fbc .CreatedAt .After (_dff ){if !_fbc .UniPDF {return _fb .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006c\u0069\u0063\u0065\u006e\u0073\u0065:\u0020\u0054\u0068\u0069\u0073\u0020\u0055\u006e\u0069\u0044\u006f\u0063\u0020k\u0065\u0079\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069d \u0066\u006f\u0072\u0020\u0055\u006e\u0069\u0050\u0044\u0046");};};return nil ;};func (_gd defaultStateHolder )loadState (_cfg string )(reportState ,error ){_dge :=_ed .Getenv ("\u0048\u004f\u004d\u0045");if len (_dge )==0{return reportState {},_de .New ("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");};_feaa :=_a .Join (_dge ,"\u002eu\u006e\u0069\u0064\u006f\u0063");_fbe :=_ed .MkdirAll (_feaa ,0777);if _fbe !=nil {return reportState {},_fbe ;};if len (_cfg )< 20{return reportState {},_de .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079");};_ecb :=[]byte (_cfg );_efd :=_ada .Sum512_256 (_ecb [:20]);_cde :=_fg .EncodeToString (_efd [:]);_cbc :=_a .Join (_feaa ,_cde );_edd ,_fbe :=_ag .ReadFile (_cbc );if _fbe !=nil {if _ed .IsNotExist (_fbe ){return reportState {},nil ;};_ga .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_fbe );return reportState {},_de .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061");};const _cag ="\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057";_edd ,_fbe =_aaa ([]byte (_cag ),_edd );if _fbe !=nil {return reportState {},_fbe ;};var _bff reportState ;_fbe =_adf .Unmarshal (_edd ,&_bff );if _fbe !=nil {_ga .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0049\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061\u003a\u0020\u0025\u0076",_fbe );return reportState {},_de .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0061");};return _bff ,nil ;};func _ffcc (_gbaa []byte )(_g .Reader ,error ){_fcbf :=new (_ef .Buffer );_gad :=_d .NewWriter (_fcbf );_gad .Write (_gbaa );_gfg :=_gad .Close ();if _gfg !=nil {return nil ,_gfg ;};return _fcbf ,nil ;};func (_cd *LicenseKey )isExpired ()bool {return _cd .getExpiryDateToCompare ().After (*_cd .ExpiresAt )};func (_dca *LicenseKey )TypeToString ()string {if _dca ._dd {return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e";};if _dca .Tier ==LicenseTierUnlicensed {return "\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";};if _dca .Tier ==LicenseTierCommunity {return "\u0041\u0047PL\u0076\u0033\u0020O\u0070\u0065\u006e\u0020Sou\u0072ce\u0020\u0043\u006f\u006d\u006d\u0075\u006eit\u0079\u0020\u004c\u0069\u0063\u0065\u006es\u0065";};if _dca .Tier ==LicenseTierIndividual ||_dca .Tier =="\u0069\u006e\u0064i\u0065"{return "\u0043\u006f\u006dm\u0065\u0072\u0063\u0069a\u006c\u0020\u004c\u0069\u0063\u0065\u006es\u0065\u0020\u002d\u0020\u0049\u006e\u0064\u0069\u0076\u0069\u0064\u0075\u0061\u006c";};return "\u0043\u006fm\u006d\u0065\u0072\u0063\u0069\u0061\u006c\u0020\u004c\u0069\u0063\u0065\u006e\u0073\u0065\u0020\u002d\u0020\u0042\u0075\u0073\u0069ne\u0073\u0073";};type stateLoader interface{loadState (_gage string )(reportState ,error );updateState (_abd ,_adbe ,_bf string ,_fgd int ,_ffbb bool ,_ccee int ,_bca int ,_eae _dc .Time ,_gafc map[string ]int )error ;};type defaultStateHolder struct{};func GetMeteredState ()(MeteredStatus ,error ){if _aea ==nil {return MeteredStatus {},_de .New ("\u006c\u0069\u0063\u0065ns\u0065\u0020\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");};if !_aea ._dd ||len (_aea ._eaa )==0{return MeteredStatus {},_de .New ("\u0061p\u0069 \u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");};_dba ,_cdd :=_beg .loadState (_aea ._eaa );if _cdd !=nil {_ga .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_cdd );return MeteredStatus {},_cdd ;};if _dba .Docs > 0{_gaf :=_cad ("","",true );if _gaf !=nil {return MeteredStatus {},_gaf ;};};_geb .Lock ();defer _geb .Unlock ();_ccc :=_ffc ();_ccc ._aeb =_aea ._eaa ;_defa ,_cdd :=_ccc .getStatus ();if _cdd !=nil {return MeteredStatus {},_cdd ;};if !_defa .Valid {return MeteredStatus {},_de .New ("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064");};_ffcd :=MeteredStatus {OK :true ,Credits :_defa .OrgCredits ,Used :_defa .OrgUsed };return _ffcd ,nil ;};func _adc ()(_gf .IP ,error ){_dgf ,_deg :=_gf .Dial ("\u0075\u0064\u0070","\u0038\u002e\u0038\u002e\u0038\u002e\u0038\u003a\u0038\u0030");if _deg !=nil {return nil ,_deg ;};defer _dgf .Close ();_cfe :=_dgf .LocalAddr ().(*_gf .UDPAddr );return _cfe .IP ,nil ;};var _ecf map[string ]struct{};type meteredStatusResp struct{Valid bool `json:"valid"`;OrgCredits int64 `json:"org_credits"`;OrgUsed int64 `json:"org_used"`;OrgRemaining int64 `json:"org_remaining"`;};type meteredUsageCheckinForm struct{Instance string `json:"inst"`;Next string `json:"next"`;UsageNumber int `json:"usage_number"`;NumFailed int64 `json:"num_failed"`;Hostname string `json:"hostname"`;LocalIP string `json:"local_ip"`;MacAddress string `json:"mac_address"`;Package string `json:"package"`;PackageVersion string `json:"package_version"`;Usage map[string ]int `json:"u"`;};var _eac map[string ]int ;func SetMeteredKey (apiKey string )error {_gcf :=_ffc ();_gcf ._aeb =apiKey ;_gac ,_dgg :=_gcf .getStatus ();if _dgg !=nil {return _dgg ;};if !_gac .Valid {return _de .New ("\u006b\u0065\u0079\u0020\u006e\u006f\u0074\u0020\u0076\u0061\u006c\u0069\u0064");};_cfc :=&LicenseKey {_dd :true ,_eaa :apiKey };_aea =_cfc ;return nil ;};func (_fea *LicenseKey )IsLicensed ()bool {return _fea .Tier !=LicenseTierUnlicensed ||_fea ._dd };type reportState struct{Instance string `json:"inst"`;Next string `json:"n"`;Docs int64 `json:"d"`;NumErrors int64 `json:"e"`;LimitDocs bool `json:"ld"`;RemainingDocs int64 `json:"rd"`;LastReported _dc .Time `json:"lr"`;LastWritten _dc .Time `json:"lw"`;Usage map[string ]int `json:"u"`;};var _beg stateLoader =defaultStateHolder {};type MeteredStatus struct{OK bool ;Credits int64 ;Used int64 ;};func (_dffa *LicenseKey )getExpiryDateToCompare ()_dc .Time {if _dffa .Trial {return _dc .Now ().UTC ();};return _ga .ReleasedAt ;};var _gga =_dc .Date (2020,1,1,0,0,0,0,_dc .UTC );var _aea =MakeUnlicensedKey ();const _abca ="\u0055\u004e\u0049\u0050DF\u005f\u004c\u0049\u0043\u0045\u004e\u0053\u0045\u005f\u0050\u0041\u0054\u0048";func MakeUnlicensedKey ()*LicenseKey {_fgg :=LicenseKey {};_fgg .CustomerName ="\u0055\u006e\u006c\u0069\u0063\u0065\u006e\u0073\u0065\u0064";_fgg .Tier =LicenseTierUnlicensed ;_fgg .CreatedAt =_dc .Now ().UTC ();_fgg .CreatedAtInt =_fgg .CreatedAt .Unix ();return &_fgg ;};func _efb (_eg string ,_dg []byte )(string ,error ){_ba ,_ :=_ad .Decode ([]byte (_eg ));if _ba ==nil {return "",_fb .Errorf ("\u0050\u0072\u0069\u0076\u004b\u0065\u0079\u0020\u0066a\u0069\u006c\u0065\u0064");};_ge ,_gg :=_cc .ParsePKCS1PrivateKey (_ba .Bytes );if _gg !=nil {return "",_gg ;};_dbb :=_ada .New ();_dbb .Write (_dg );_gge :=_dbb .Sum (nil );_dcg ,_gg :=_def .SignPKCS1v15 (_ab .Reader ,_ge ,_e .SHA512 ,_gge );if _gg !=nil {return "",_gg ;};_bg :=_bc .StdEncoding .EncodeToString (_dg );_bg +="\u000a\u002b\u000a";_bg +=_bc .StdEncoding .EncodeToString (_dcg );return _bg ,nil ;};func (_bcd *LicenseKey )ToString ()string {if _bcd ._dd {return "M\u0065t\u0065\u0072\u0065\u0064\u0020\u0073\u0075\u0062s\u0063\u0072\u0069\u0070ti\u006f\u006e";};_dea :=_fb .Sprintf ("\u004ci\u0063e\u006e\u0073\u0065\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_bcd .LicenseId );_dea +=_fb .Sprintf ("\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u0049\u0064\u003a\u0020\u0025\u0073\u000a",_bcd .CustomerId );_dea +=_fb .Sprintf ("\u0043u\u0073t\u006f\u006d\u0065\u0072\u0020N\u0061\u006de\u003a\u0020\u0025\u0073\u000a",_bcd .CustomerName );_dea +=_fb .Sprintf ("\u0054i\u0065\u0072\u003a\u0020\u0025\u0073\n",_bcd .Tier );_dea +=_fb .Sprintf ("\u0043r\u0065a\u0074\u0065\u0064\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_ga .UtcTimeFormat (_bcd .CreatedAt ));if _bcd .ExpiresAt ==nil {_dea +="\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041t\u003a\u0020N\u0065\u0076\u0065\u0072\u000a";}else {_dea +=_fb .Sprintf ("\u0045x\u0070i\u0072\u0065\u0073\u0020\u0041\u0074\u003a\u0020\u0025\u0073\u000a",_ga .UtcTimeFormat (*_bcd .ExpiresAt ));};_dea +=_fb .Sprintf ("\u0043\u0072\u0065\u0061\u0074\u006f\u0072\u003a\u0020\u0025\u0073\u0020<\u0025\u0073\u003e\u000a",_bcd .CreatorName ,_bcd .CreatorEmail );return _dea ;};func GetLicenseKey ()*LicenseKey {if _aea ==nil {return nil ;};_acd :=*_aea ;return &_acd ;};type meteredStatusForm struct{};func Track (docKey string ,useKey string )error {return _cad (docKey ,useKey ,false )};type meteredUsageCheckinResp struct{Instance string `json:"inst"`;Next string `json:"next"`;Success bool `json:"success"`;Message string `json:"message"`;RemainingDocs int `json:"rd"`;LimitDocs bool `json:"ld"`;};var _dff =_dc .Date (2019,6,6,0,0,0,0,_dc .UTC );func (_abed *meteredClient )checkinUsage (_abc meteredUsageCheckinForm )(meteredUsageCheckinResp ,error ){_abc .Package ="\u0075\u006e\u0069\u0070\u0064\u0066";_abc .PackageVersion =_ga .Version ;var _gbc meteredUsageCheckinResp ;_afc :=_abed ._ddd +"\u002f\u006d\u0065\u0074er\u0065\u0064\u002f\u0075\u0073\u0061\u0067\u0065\u005f\u0063\u0068\u0065\u0063\u006bi\u006e";_dda ,_ca :=_adf .Marshal (_abc );if _ca !=nil {return _gbc ,_ca ;};_aa ,_ca :=_ffcc (_dda );if _ca !=nil {return _gbc ,_ca ;};_fgc ,_ca :=_be .NewRequest ("\u0050\u004f\u0053\u0054",_afc ,_aa );if _ca !=nil {return _gbc ,_ca ;};_fgc .Header .Add ("\u0043\u006f\u006et\u0065\u006e\u0074\u002d\u0054\u0079\u0070\u0065","\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e");_fgc .Header .Add ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_fgc .Header .Add ("\u0041c\u0063e\u0070\u0074\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067","\u0067\u007a\u0069\u0070");_fgc .Header .Add ("\u0058-\u0041\u0050\u0049\u002d\u004b\u0045Y",_abed ._aeb );_egf ,_ca :=_abed ._ggea .Do (_fgc );if _ca !=nil {return _gbc ,_ca ;};defer _egf .Body .Close ();if _egf .StatusCode !=200{return _gbc ,_fb .Errorf ("\u0066\u0061i\u006c\u0065\u0064\u0020t\u006f\u0020c\u0068\u0065\u0063\u006b\u0069\u006e\u002c\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u0020\u0069s\u003a\u0020\u0025\u0064",_egf .StatusCode );};_dfc ,_ca :=_fggb (_egf );if _ca !=nil {return _gbc ,_ca ;};_ca =_adf .Unmarshal (_dfc ,&_gbc );if _ca !=nil {return _gbc ,_ca ;};return _gbc ,nil ;};func SetLicenseKey (content string ,customerName string )error {_ffg ,_efcg :=_facf (content );if _efcg !=nil {_ga .Log .Error ("\u004c\u0069c\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0065\u0072\u0072\u006f\u0072: \u0025\u0076",_efcg );return _efcg ;};if !_efg .EqualFold (_ffg .CustomerName ,customerName ){_ga .Log .Error ("L\u0069ce\u006es\u0065 \u0063\u006f\u0064\u0065\u0020i\u0073\u0073\u0075e\u0020\u002d\u0020\u0043\u0075s\u0074\u006f\u006de\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006da\u0074\u0063\u0068, e\u0078\u0070\u0065\u0063\u0074\u0065d\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067o\u0074 \u0027\u0025\u0073\u0027",customerName ,_ffg .CustomerName );return _fb .Errorf ("\u0063\u0075\u0073\u0074\u006fm\u0065\u0072\u0020\u006e\u0061\u006d\u0065\u0020\u006d\u0069\u0073\u006d\u0061t\u0063\u0068\u002c\u0020\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0027\u0025\u0073\u0027\u002c\u0020\u0062\u0075\u0074\u0020\u0067\u006f\u0074\u0020\u0027\u0025\u0073'",customerName ,_ffg .CustomerName );};_efcg =_ffg .Validate ();if _efcg !=nil {_ga .Log .Error ("\u004c\u0069\u0063\u0065\u006e\u0073e\u0020\u0063\u006f\u0064\u0065\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074i\u006f\u006e\u0020\u0065\u0072\u0072\u006fr\u003a\u0020\u0025\u0076",_efcg );return _efcg ;};_aea =&_ffg ;return nil ;};func _ee (_ffbe string ,_af string ,_fac string )(string ,error ){_cbe :=_efg .Index (_fac ,_ffbe );if _cbe ==-1{return "",_fb .Errorf ("\u0068\u0065a\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_dcd :=_efg .Index (_fac ,_af );if _dcd ==-1{return "",_fb .Errorf ("\u0066\u006fo\u0074\u0065\u0072 \u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_aba :=_cbe +len (_ffbe )+1;return _fac [_aba :_dcd -1],nil ;};func _aaa (_bfgb ,_cfb []byte )([]byte ,error ){_degc :=make ([]byte ,_bc .URLEncoding .DecodedLen (len (_cfb )));_adgf ,_bffc :=_bc .URLEncoding .Decode (_degc ,_cfb );if _bffc !=nil {return nil ,_bffc ;};_degc =_degc [:_adgf ];_acg ,_bffc :=_db .NewCipher (_bfgb );if _bffc !=nil {return nil ,_bffc ;};if len (_degc )< _db .BlockSize {return nil ,_de .New ("c\u0069p\u0068\u0065\u0072\u0074\u0065\u0078\u0074\u0020t\u006f\u006f\u0020\u0073ho\u0072\u0074");};_cfbf :=_degc [:_db .BlockSize ];_degc =_degc [_db .BlockSize :];_eddc :=_fa .NewCFBDecrypter (_acg ,_cfbf );_eddc .XORKeyStream (_degc ,_degc );return _degc ,nil ;};const _eaag ="\u000a\u002d\u002d\u002d\u002d\u002d\u0042\u0045\u0047\u0049\u004e \u0050\u0055\u0042\u004c\u0049\u0043\u0020\u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\u000a\u004d\u0049I\u0042\u0049\u006a\u0041NB\u0067\u006b\u0071\u0068\u006b\u0069G\u0039\u0077\u0030\u0042\u0041\u0051\u0045\u0046A\u0041\u004f\u0043\u0041\u0051\u0038\u0041\u004d\u0049\u0049\u0042\u0043\u0067\u004b\u0043\u0041\u0051\u0045A\u006dF\u0055\u0069\u0079\u0064\u0037\u0062\u0035\u0058\u006a\u0070\u006b\u0050\u0035\u0052\u0061\u0070\u0034\u0077\u000a\u0044\u0063\u0031d\u0079\u007a\u0049\u0051\u0034\u004c\u0065\u006b\u0078\u0072\u0076\u0079\u0074\u006e\u0045\u004d\u0070\u004e\u0055\u0062\u006f\u0036i\u0041\u0037\u0034\u0056\u0038\u0072\u0075\u005a\u004f\u0076\u0072\u0053\u0063\u0073\u0066\u0032\u0051\u0065\u004e9\u002f\u0071r\u0055\u0047\u0038\u0071\u0045\u0062\u0055\u0057\u0064\u006f\u0045\u0059\u0071+\u000a\u006f\u0074\u0046\u004e\u0041\u0046N\u0078\u006c\u0047\u0062\u0078\u0062\u0044\u0048\u0063\u0064\u0047\u0056\u0061\u004d\u0030\u004f\u0058\u0064\u0058g\u0044y\u004c5\u0061\u0049\u0045\u0061\u0067\u004c\u0030\u0063\u0035\u0070\u0077\u006a\u0049\u0064\u0050G\u0049\u006e\u0034\u0036\u0066\u0037\u0038\u0065\u004d\u004a\u002b\u004a\u006b\u0064\u0063\u0070\u0044\n\u0044\u004a\u0061\u0071\u0059\u0058d\u0072\u007a5\u004b\u0065\u0073\u0068\u006aS\u0069\u0049\u0061\u0061\u0037\u006d\u0065\u006e\u0042\u0049\u0041\u0058\u0053\u0034\u0055\u0046\u0078N\u0066H\u0068\u004e\u0030\u0048\u0043\u0059\u005a\u0059\u0071\u0051\u0047\u0037\u0062K+\u0073\u0035\u0072R\u0048\u006f\u006e\u0079\u0064\u004eW\u0045\u0047\u000a\u0048\u0038M\u0079\u0076\u00722\u0070\u0079\u0061\u0032K\u0072\u004d\u0075m\u0066\u006d\u0041\u0078\u0055\u0042\u0036\u0066\u0065\u006e\u0043\u002f4\u004f\u0030\u0057\u00728\u0067\u0066\u0050\u004f\u0055\u0038R\u0069\u0074\u006d\u0062\u0044\u0076\u0051\u0050\u0049\u0052\u0058\u004fL\u0034\u0076\u0054B\u0072\u0042\u0064\u0062a\u0041\u000a9\u006e\u0077\u004e\u0050\u002b\u0069\u002f\u002f\u0032\u0030\u004d\u00542\u0062\u0078\u006d\u0065\u0057\u0042\u002b\u0067\u0070\u0063\u0045\u0068G\u0070\u0058\u005a7\u0033\u0033\u0061\u007a\u0051\u0078\u0072\u0043\u0033\u004a\u0034\u0076\u0033C\u005a\u006d\u0045\u004eS\u0074\u0044\u004b\u002f\u004b\u0044\u0053\u0050\u004b\u0055\u0047\u0066\u00756\u000a\u0066\u0077I\u0044\u0041\u0051\u0041\u0042\u000a\u002d\u002d\u002d\u002d\u002dE\u004e\u0044\u0020\u0050\u0055\u0042\u004c\u0049\u0043 \u004b\u0045Y\u002d\u002d\u002d\u002d\u002d\n";func _ffc ()*meteredClient {_cce :=meteredClient {_ddd :"\u0068t\u0074\u0070\u0073\u003a\u002f\u002f\u0075\u006e\u0069\u0063\u0072m\u002e\u0075\u006e\u0069\u0064\u006f\u0063\u002e\u0069\u006f",_ggea :&_be .Client {Timeout :30*_dc .Second }};if _dbe :=_ed .Getenv ("\u0055N\u0049\u0044\u004f\u0043_\u004c\u0049\u0043\u0045\u004eS\u0045_\u0053E\u0052\u0056\u0045\u0052\u005f\u0055\u0052L");_efg .HasPrefix (_dbe ,"\u0068\u0074\u0074\u0070"){_cce ._ddd =_dbe ;};return &_cce ;};func TrackUse (useKey string ){if _aea ==nil {return ;};if !_aea ._dd ||len (_aea ._eaa )==0{return ;};if len (useKey )==0{return ;};_geb .Lock ();defer _geb .Unlock ();if _eac ==nil {_eac =map[string ]int {};};_eac [useKey ]++;};type meteredClient struct{_ddd string ;_aeb string ;_ggea *_be .Client ;};func _cadc (_faeb *_be .Response )(_g .ReadCloser ,error ){var _ggd error ;var _fece _g .ReadCloser ;switch _efg .ToLower (_faeb .Header .Get ("\u0043\u006fn\u0074\u0065\u006et\u002d\u0045\u006e\u0063\u006f\u0064\u0069\u006e\u0067")){case "\u0067\u007a\u0069\u0070":_fece ,_ggd =_d .NewReader (_faeb .Body );if _ggd !=nil {return _fece ,_ggd ;};defer _fece .Close ();default:_fece =_faeb .Body ;};return _fece ,nil ;};const _feg ="U\u004eI\u0050\u0044\u0046\u005f\u0043\u0055\u0053\u0054O\u004d\u0045\u0052\u005fNA\u004d\u0045";func (_fec defaultStateHolder )updateState (_edg ,_eefe ,_bbf string ,_deec int ,_ace bool ,_eag int ,_dfa int ,_dce _dc .Time ,_fba map[string ]int )error {_dcf :=_ed .Getenv ("\u0048\u004f\u004d\u0045");if len (_dcf )==0{return _de .New ("\u0068\u006fm\u0065\u0020\u0064i\u0072\u0020\u006e\u006f\u0074\u0020\u0073\u0065\u0074");};_cg :=_a .Join (_dcf ,"\u002eu\u006e\u0069\u0064\u006f\u0063");_afd :=_ed .MkdirAll (_cg ,0777);if _afd !=nil {return _afd ;};if len (_edg )< 20{return _de .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006b\u0065\u0079");};_bab :=[]byte (_edg );_bfa :=_ada .Sum512_256 (_bab [:20]);_ffd :=_fg .EncodeToString (_bfa [:]);_aad :=_a .Join (_cg ,_ffd );var _ec reportState ;_ec .Docs =int64 (_deec );_ec .NumErrors =int64 (_dfa );_ec .LimitDocs =_ace ;_ec .RemainingDocs =int64 (_eag );_ec .LastWritten =_dc .Now ().UTC ();_ec .LastReported =_dce ;_ec .Instance =_eefe ;_ec .Next =_bbf ;_ec .Usage =_fba ;_afdg ,_afd :=_adf .Marshal (_ec );if _afd !=nil {return _afd ;};const _dcfb ="\u0068\u00619\u004e\u004b\u0038]\u0052\u0062\u004c\u002a\u006d\u0034\u004c\u004b\u0057";_afdg ,_afd =_ggfc ([]byte (_dcfb ),_afdg );if _afd !=nil {return _afd ;};_afd =_ag .WriteFile (_aad ,_afdg ,0600);if _afd !=nil {return _afd ;};return nil ;};func init (){_abf :=_ed .Getenv (_abca );_aee :=_ed .Getenv (_feg );if len (_abf )==0||len (_aee )==0{return ;};_daa ,_ffcb :=_ag .ReadFile (_abf );if _ffcb !=nil {_ga .Log .Error ("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0072\u0065ad \u006cic\u0065\u006e\u0073\u0065\u0020\u0063\u006fde\u0020\u0066\u0069\u006c\u0065\u003a\u0020%\u0076",_ffcb );return ;};_ffcb =SetLicenseKey (string (_daa ),_aee );if _ffcb !=nil {_ga .Log .Error ("\u0055\u006e\u0061b\u006c\u0065\u0020\u0074o\u0020\u006c\u006f\u0061\u0064\u0020\u006ci\u0063\u0065\u006e\u0073\u0065\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0076",_ffcb );return ;};};func _cgf ()([]string ,[]string ,error ){_aff ,_fcb :=_gf .Interfaces ();if _fcb !=nil {return nil ,nil ,_fcb ;};var _dcfbd []string ;var _cfd []string ;for _ ,_gff :=range _aff {if _gff .Flags &_gf .FlagUp ==0||_ef .Equal (_gff .HardwareAddr ,nil ){continue ;};if _gff .HardwareAddr [0]&2==2{continue ;};_ggf ,_cbf :=_gff .Addrs ();if _cbf !=nil {return nil ,nil ,_cbf ;};_ecg :=0;for _ ,_eagg :=range _ggf {var _ege _gf .IP ;switch _bec :=_eagg .(type ){case *_gf .IPNet :_ege =_bec .IP ;case *_gf .IPAddr :_ege =_bec .IP ;};if _ege .IsLoopback (){continue ;};if _ege .To4 ()==nil {continue ;};_cfd =append (_cfd ,_ege .String ());_ecg ++;};_bda :=_gff .HardwareAddr .String ();if _bda !=""&&_ecg > 0{_dcfbd =append (_dcfbd ,_bda );};};return _dcfbd ,_cfd ,nil ;};func _ggfc (_bde ,_ddb []byte )([]byte ,error ){_aeea ,_agg :=_db .NewCipher (_bde );if _agg !=nil {return nil ,_agg ;};_ccdb :=make ([]byte ,_db .BlockSize +len (_ddb ));_fegd :=_ccdb [:_db .BlockSize ];if _ ,_ccdg :=_g .ReadFull (_ab .Reader ,_fegd );_ccdg !=nil {return nil ,_ccdg ;};_dfe :=_fa .NewCFBEncrypter (_aeea ,_fegd );_dfe .XORKeyStream (_ccdb [_db .BlockSize :],_ddb );_ccec :=make ([]byte ,_bc .URLEncoding .EncodedLen (len (_ccdb )));_bc .URLEncoding .Encode (_ccec ,_ccdb );return _ccec ,nil ;};var _eba =_dc .Date (2010,1,1,0,0,0,0,_dc .UTC );const (_abe ="\u002d\u002d\u002d--\u0042\u0045\u0047\u0049\u004e\u0020\u0055\u004e\u0049D\u004fC\u0020L\u0049C\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";_bb ="\u002d\u002d\u002d\u002d\u002d\u0045\u004e\u0044\u0020\u0055\u004e\u0049\u0044\u004f\u0043 \u004cI\u0043\u0045\u004e\u0053\u0045\u0020\u004b\u0045\u0059\u002d\u002d\u002d\u002d\u002d";);func _fggb (_degd *_be .Response )([]byte ,error ){var _affc []byte ;_dde ,_cgc :=_cadc (_degd );if _cgc !=nil {return _affc ,_cgc ;};return _ag .ReadAll (_dde );};func _cb (_cba string ,_ac string )([]byte ,error ){var (_fc int ;_ffb string ;);for _ ,_ffb =range []string {"\u000a\u002b\u000a","\u000d\u000a\u002b\r\u000a","\u0020\u002b\u0020"}{if _fc =_efg .Index (_ac ,_ffb );_fc !=-1{break ;};};if _fc ==-1{return nil ,_fb .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u002c \u0073i\u0067n\u0061t\u0075\u0072\u0065\u0020\u0073\u0065\u0070\u0061\u0072\u0061\u0074\u006f\u0072");};_fbf :=_ac [:_fc ];_efa :=_fc +len (_ffb );_ffa :=_ac [_efa :];if _fbf ==""||_ffa ==""{return nil ,_fb .Errorf ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0069n\u0070\u0075\u0074,\u0020\u006d\u0069\u0073\u0073\u0069\u006e\u0067\u0020or\u0069\u0067\u0069n\u0061\u006c \u006f\u0072\u0020\u0073\u0069\u0067n\u0061\u0074u\u0072\u0065");};_cbb ,_bd :=_bc .StdEncoding .DecodeString (_fbf );if _bd !=nil {return nil ,_fb .Errorf ("\u0069\u006e\u0076\u0061li\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0072\u0069\u0067\u0069\u006ea\u006c");};_gb ,_bd :=_bc .StdEncoding .DecodeString (_ffa );if _bd !=nil {return nil ,_fb .Errorf ("\u0069\u006e\u0076al\u0069\u0064\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");};_df ,_ :=_ad .Decode ([]byte (_cba ));if _df ==nil {return nil ,_fb .Errorf ("\u0050\u0075\u0062\u004b\u0065\u0079\u0020\u0066\u0061\u0069\u006c\u0065\u0064");};_bdf ,_bd :=_cc .ParsePKIXPublicKey (_df .Bytes );if _bd !=nil {return nil ,_bd ;};_gc :=_bdf .(*_def .PublicKey );if _gc ==nil {return nil ,_fb .Errorf ("\u0050u\u0062\u004b\u0065\u0079\u0020\u0063\u006f\u006e\u0076\u0065\u0072s\u0069\u006f\u006e\u0020\u0066\u0061\u0069\u006c\u0065\u0064");};_ae :=_ada .New ();_ae .Write (_cbb );_gag :=_ae .Sum (nil );_bd =_def .VerifyPKCS1v15 (_gc ,_e .SHA512 ,_gag ,_gb );if _bd !=nil {return nil ,_bd ;};return _cbb ,nil ;};