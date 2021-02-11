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

package mmr ;import (_b "errors";_g "fmt";_c "github.com/unidoc/unipdf/v3/common";_fe "github.com/unidoc/unipdf/v3/internal/bitwise";_gc "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_f "io";);func (_dad *Decoder )initTables ()(_aa error ){if _dad ._dea ==nil {_dad ._dea ,_aa =_dad .createLittleEndianTable (_dc );if _aa !=nil {return ;};_dad ._ffb ,_aa =_dad .createLittleEndianTable (_gg );if _aa !=nil {return ;};_dad ._ceb ,_aa =_dad .createLittleEndianTable (_fb );if _aa !=nil {return ;};};return nil ;};func (_ffc *Decoder )fillBitmap (_befg *_gc .Bitmap ,_dba int ,_bbd []int ,_gbe int )error {var _cfc byte ;_fef :=0;_acb :=_befg .GetByteIndex (_fef ,_dba );for _dbc :=0;_dbc < _gbe ;_dbc ++{_dbf :=byte (1);_dfa :=_bbd [_dbc ];if (_dbc &1)==0{_dbf =0;};for _fef < _dfa {_cfc =(_cfc <<1)|_dbf ;_fef ++;if (_fef &7)==0{if _bga :=_befg .SetByte (_acb ,_cfc );_bga !=nil {return _bga ;};_acb ++;_cfc =0;};};};if (_fef &7)!=0{_cfc <<=uint (8-(_fef &7));if _eee :=_befg .SetByte (_acb ,_cfc );_eee !=nil {return _eee ;};};return nil ;};type mmrCode int ;const (EOF =-3;_bg =-2;EOL =-1;_ab =8;_ff =(1<<_ab )-1;_gcg =5;_dg =(1<<_gcg )-1;);func (_cdc *Decoder )createLittleEndianTable (_cda [][3]int )([]*code ,error ){_aca :=make ([]*code ,_ff +1);for _ef :=0;_ef < len (_cda );_ef ++{_ad :=_a (_cda [_ef ]);if _ad ._cb <=_ab {_fda :=_ab -_ad ._cb ;_fdcd :=_ad ._bd <<uint (_fda );for _ba :=(1<<uint (_fda ))-1;_ba >=0;_ba --{_ebd :=_fdcd |_ba ;_aca [_ebd ]=_ad ;};}else {_ae :=_ad ._bd >>uint (_ad ._cb -_ab );if _aca [_ae ]==nil {var _cge =_a ([3]int {});_cge ._e =make ([]*code ,_dg +1);_aca [_ae ]=_cge ;};if _ad ._cb <=_ab +_gcg {_dfe :=_ab +_gcg -_ad ._cb ;_bef :=(_ad ._bd <<uint (_dfe ))&_dg ;_aca [_ae ]._eg =true ;for _db :=(1<<uint (_dfe ))-1;_db >=0;_db --{_aca [_ae ]._e [_bef |_db ]=_ad ;};}else {return nil ,_b .New ("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072");};};};return _aca ,nil ;};func (_bgg *runData )fillBuffer (_gbb int )error {_bgg ._ec =_gbb ;_ ,_cdag :=_bgg ._efd .Seek (int64 (_gbb ),_f .SeekStart );if _cdag !=nil {if _cdag ==_f .EOF {_c .Log .Debug ("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046");_bgg ._ggc =-1;}else {return _cdag ;};};if _cdag ==nil {_bgg ._ggc ,_cdag =_bgg ._efd .Read (_bgg ._edd );if _cdag !=nil {if _cdag ==_f .EOF {_c .Log .Trace ("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046");_bgg ._ggc =-1;}else {return _cdag ;};};};if _bgg ._ggc > -1&&_bgg ._ggc < 3{for _bgg ._ggc < 3{_cfb ,_bfc :=_bgg ._efd .ReadByte ();if _bfc !=nil {if _bfc ==_f .EOF {_bgg ._edd [_bgg ._ggc ]=0;}else {return _bfc ;};}else {_bgg ._edd [_bgg ._ggc ]=_cfb &0xFF;};_bgg ._ggc ++;};};_bgg ._ggc -=3;if _bgg ._ggc < 0{_bgg ._edd =make ([]byte ,len (_bgg ._edd ));_bgg ._ggc =len (_bgg ._edd )-3;};return nil ;};func (_cea *Decoder )uncompress1d (_abe *runData ,_daf []int ,_edg int )(int ,error ){var (_bbc =true ;_fgd int ;_bbe *code ;_ffcb int ;_fga error ;);_gf :for _fgd < _edg {_fab :for {if _bbc {_bbe ,_fga =_abe .uncompressGetCode (_cea ._dea );if _fga !=nil {return 0,_fga ;};}else {_bbe ,_fga =_abe .uncompressGetCode (_cea ._ffb );if _fga !=nil {return 0,_fga ;};};_abe ._bed +=_bbe ._cb ;if _bbe ._gd < 0{break _gf ;};_fgd +=_bbe ._gd ;if _bbe ._gd < 64{_bbc =!_bbc ;_daf [_ffcb ]=_fgd ;_ffcb ++;break _fab ;};};};if _daf [_ffcb ]!=_edg {_daf [_ffcb ]=_edg ;};_dgb :=EOL ;if _bbe !=nil &&_bbe ._gd !=EOL {_dgb =_ffcb ;};return _dgb ,nil ;};type Decoder struct{_abf ,_acd int ;_abd *runData ;_dea []*code ;_ffb []*code ;_ceb []*code ;};func (_dfg *code )String ()string {return _g .Sprintf ("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_dfg ._cb ,_dfg ._bd ,_dfg ._gd );};func _beg (_eea *_fe .SubstreamReader )(*runData ,error ){_dee :=&runData {_efd :_eea ,_bed :0,_edf :1};_cfce :=_fg (_bda (_ga ,int (_eea .Length ())),_cbe );_dee ._edd =make ([]byte ,_cfce );if _efa :=_dee .fillBuffer (0);_efa !=nil {if _efa ==_f .EOF {_dee ._edd =make ([]byte ,10);_c .Log .Debug ("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076",_efa );}else {return nil ,_efa ;};};return _dee ,nil ;};func (_dfd *runData )align (){_dfd ._bed =((_dfd ._bed +7)>>3)<<3};func New (r _fe .StreamReader ,width ,height int ,dataOffset ,dataLength int64 )(*Decoder ,error ){_fdc :=&Decoder {_abf :width ,_acd :height };_be ,_cf :=_fe .NewSubstreamReader (r ,uint64 (dataOffset ),uint64 (dataLength ));if _cf !=nil {return nil ,_cf ;};_bea ,_cf :=_beg (_be );if _cf !=nil {return nil ,_cf ;};_fdc ._abd =_bea ;if _eb :=_fdc .initTables ();_eb !=nil {return nil ,_eb ;};return _fdc ,nil ;};func _a (_df [3]int )*code {return &code {_cb :_df [0],_bd :_df [1],_gd :_df [2]}};func _fg (_ed ,_fd int )int {if _ed > _fd {return _fd ;};return _ed ;};func (_ag *Decoder )UncompressMMR ()(_edc *_gc .Bitmap ,_cde error ){_edc =_gc .New (_ag ._abf ,_ag ._acd );_cbf :=make ([]int ,_edc .Width +5);_gdb :=make ([]int ,_edc .Width +5);_gdb [0]=_edc .Width ;_fff :=1;var _dgc int ;for _cg :=0;_cg < _edc .Height ;_cg ++{_dgc ,_cde =_ag .uncompress2d (_ag ._abd ,_gdb ,_fff ,_cbf ,_edc .Width );if _cde !=nil {return nil ,_cde ;};if _dgc ==EOF {break ;};if _dgc > 0{_cde =_ag .fillBitmap (_edc ,_cg ,_cbf ,_dgc );if _cde !=nil {return nil ,_cde ;};};_gdb ,_cbf =_cbf ,_gdb ;_fff =_dgc ;};if _cde =_ag .detectAndSkipEOL ();_cde !=nil {return nil ,_cde ;};_ag ._abd .align ();return _edc ,nil ;};const (_cbe int =1024<<7;_ga int =3;_egb uint =24;);func (_fdg *runData )uncompressGetCode (_ggb []*code )(*code ,error ){return _fdg .uncompressGetCodeLittleEndian (_ggb );};type runData struct{_efd *_fe .SubstreamReader ;_bed int ;_edf int ;_fdf int ;_edd []byte ;_ec int ;_ggc int ;};func _bda (_de ,_cd int )int {if _de < _cd {return _cd ;};return _de ;};const (_ee mmrCode =iota ;_ac ;_ce ;_bb ;_da ;_fea ;_gb ;_ca ;_cdf ;_egg ;_cc ;);func (_cbb *runData )uncompressGetCodeLittleEndian (_faa []*code )(*code ,error ){_dab ,_ada :=_cbb .uncompressGetNextCodeLittleEndian ();if _ada !=nil {_c .Log .Debug ("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076",_ada );return nil ,_ada ;};_dab &=0xffffff;_befgd :=_dab >>(_egb -_ab );_fdcc :=_faa [_befgd ];if _fdcc !=nil &&_fdcc ._eg {_befgd =(_dab >>(_egb -_ab -_gcg ))&_dg ;_fdcc =_fdcc ._e [_befgd ];};return _fdcc ,nil ;};type code struct{_cb int ;_bd int ;_gd int ;_e []*code ;_eg bool ;};func (_bc *Decoder )uncompress2d (_dag *runData ,_dec []int ,_fc int ,_cfg []int ,_bf int )(int ,error ){var (_fdb int ;_ced int ;_bge int ;_dfaa =true ;_cbc error ;_aef *code ;);_dec [_fc ]=_bf ;_dec [_fc +1]=_bf ;_dec [_fc +2]=_bf +1;_dec [_fc +3]=_bf +1;_af :for _bge < _bf {_aef ,_cbc =_dag .uncompressGetCode (_bc ._ceb );if _cbc !=nil {return EOL ,nil ;};if _aef ==nil {_dag ._bed ++;break _af ;};_dag ._bed +=_aef ._cb ;switch mmrCode (_aef ._gd ){case _ce :_bge =_dec [_fdb ];case _bb :_bge =_dec [_fdb ]+1;case _gb :_bge =_dec [_fdb ]-1;case _ac :for {var _aba []*code ;if _dfaa {_aba =_bc ._dea ;}else {_aba =_bc ._ffb ;};_aef ,_cbc =_dag .uncompressGetCode (_aba );if _cbc !=nil {return 0,_cbc ;};if _aef ==nil {break _af ;};_dag ._bed +=_aef ._cb ;if _aef ._gd < 64{if _aef ._gd < 0{_cfg [_ced ]=_bge ;_ced ++;_aef =nil ;break _af ;};_bge +=_aef ._gd ;_cfg [_ced ]=_bge ;_ced ++;break ;};_bge +=_aef ._gd ;};_bgab :=_bge ;_bde :for {var _fffb []*code ;if !_dfaa {_fffb =_bc ._dea ;}else {_fffb =_bc ._ffb ;};_aef ,_cbc =_dag .uncompressGetCode (_fffb );if _cbc !=nil {return 0,_cbc ;};if _aef ==nil {break _af ;};_dag ._bed +=_aef ._cb ;if _aef ._gd < 64{if _aef ._gd < 0{_cfg [_ced ]=_bge ;_ced ++;break _af ;};_bge +=_aef ._gd ;if _bge < _bf ||_bge !=_bgab {_cfg [_ced ]=_bge ;_ced ++;};break _bde ;};_bge +=_aef ._gd ;};for _bge < _bf &&_dec [_fdb ]<=_bge {_fdb +=2;};continue _af ;case _ee :_fdb ++;_bge =_dec [_fdb ];_fdb ++;continue _af ;case _da :_bge =_dec [_fdb ]+2;case _ca :_bge =_dec [_fdb ]-2;case _fea :_bge =_dec [_fdb ]+3;case _cdf :_bge =_dec [_fdb ]-3;default:if _dag ._bed ==12&&_aef ._gd ==EOL {_dag ._bed =0;if _ ,_cbc =_bc .uncompress1d (_dag ,_dec ,_bf );_cbc !=nil {return 0,_cbc ;};_dag ._bed ++;if _ ,_cbc =_bc .uncompress1d (_dag ,_cfg ,_bf );_cbc !=nil {return 0,_cbc ;};_dbff ,_ded :=_bc .uncompress1d (_dag ,_dec ,_bf );if _ded !=nil {return EOF ,_ded ;};_dag ._bed ++;return _dbff ,nil ;};_bge =_bf ;continue _af ;};if _bge <=_bf {_dfaa =!_dfaa ;_cfg [_ced ]=_bge ;_ced ++;if _fdb > 0{_fdb --;}else {_fdb ++;};for _bge < _bf &&_dec [_fdb ]<=_bge {_fdb +=2;};};};if _cfg [_ced ]!=_bf {_cfg [_ced ]=_bf ;};if _aef ==nil {return EOL ,nil ;};return _ced ,nil ;};var (_fb =[][3]int {{4,0x1,int (_ee )},{3,0x1,int (_ac )},{1,0x1,int (_ce )},{3,0x3,int (_bb )},{6,0x3,int (_da )},{7,0x3,int (_fea )},{3,0x2,int (_gb )},{6,0x2,int (_ca )},{7,0x2,int (_cdf )},{10,0xf,int (_egg )},{12,0xf,int (_cc )},{12,0x1,int (EOL )}};_dc =[][3]int {{4,0x07,2},{4,0x08,3},{4,0x0B,4},{4,0x0C,5},{4,0x0E,6},{4,0x0F,7},{5,0x12,128},{5,0x13,8},{5,0x14,9},{5,0x1B,64},{5,0x07,10},{5,0x08,11},{6,0x17,192},{6,0x18,1664},{6,0x2A,16},{6,0x2B,17},{6,0x03,13},{6,0x34,14},{6,0x35,15},{6,0x07,1},{6,0x08,12},{7,0x13,26},{7,0x17,21},{7,0x18,28},{7,0x24,27},{7,0x27,18},{7,0x28,24},{7,0x2B,25},{7,0x03,22},{7,0x37,256},{7,0x04,23},{7,0x08,20},{7,0xC,19},{8,0x12,33},{8,0x13,34},{8,0x14,35},{8,0x15,36},{8,0x16,37},{8,0x17,38},{8,0x1A,31},{8,0x1B,32},{8,0x02,29},{8,0x24,53},{8,0x25,54},{8,0x28,39},{8,0x29,40},{8,0x2A,41},{8,0x2B,42},{8,0x2C,43},{8,0x2D,44},{8,0x03,30},{8,0x32,61},{8,0x33,62},{8,0x34,63},{8,0x35,0},{8,0x36,320},{8,0x37,384},{8,0x04,45},{8,0x4A,59},{8,0x4B,60},{8,0x5,46},{8,0x52,49},{8,0x53,50},{8,0x54,51},{8,0x55,52},{8,0x58,55},{8,0x59,56},{8,0x5A,57},{8,0x5B,58},{8,0x64,448},{8,0x65,512},{8,0x67,640},{8,0x68,576},{8,0x0A,47},{8,0x0B,48},{9,0x01,_bg },{9,0x98,1472},{9,0x99,1536},{9,0x9A,1600},{9,0x9B,1728},{9,0xCC,704},{9,0xCD,768},{9,0xD2,832},{9,0xD3,896},{9,0xD4,960},{9,0xD5,1024},{9,0xD6,1088},{9,0xD7,1152},{9,0xD8,1216},{9,0xD9,1280},{9,0xDA,1344},{9,0xDB,1408},{10,0x01,_bg },{11,0x01,_bg },{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560}};_gg =[][3]int {{2,0x02,3},{2,0x03,2},{3,0x02,1},{3,0x03,4},{4,0x02,6},{4,0x03,5},{5,0x03,7},{6,0x04,9},{6,0x05,8},{7,0x04,10},{7,0x05,11},{7,0x07,12},{8,0x04,13},{8,0x07,14},{9,0x01,_bg },{9,0x18,15},{10,0x01,_bg },{10,0x17,16},{10,0x18,17},{10,0x37,0},{10,0x08,18},{10,0x0F,64},{11,0x01,_bg },{11,0x17,24},{11,0x18,25},{11,0x28,23},{11,0x37,22},{11,0x67,19},{11,0x68,20},{11,0x6C,21},{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560},{12,0x24,52},{12,0x27,55},{12,0x28,56},{12,0x2B,59},{12,0x2C,60},{12,0x33,320},{12,0x34,384},{12,0x35,448},{12,0x37,53},{12,0x38,54},{12,0x52,50},{12,0x53,51},{12,0x54,44},{12,0x55,45},{12,0x56,46},{12,0x57,47},{12,0x58,57},{12,0x59,58},{12,0x5A,61},{12,0x5B,256},{12,0x64,48},{12,0x65,49},{12,0x66,62},{12,0x67,63},{12,0x68,30},{12,0x69,31},{12,0x6A,32},{12,0x6B,33},{12,0x6C,40},{12,0x6D,41},{12,0xC8,128},{12,0xC9,192},{12,0xCA,26},{12,0xCB,27},{12,0xCC,28},{12,0xCD,29},{12,0xD2,34},{12,0xD3,35},{12,0xD4,36},{12,0xD5,37},{12,0xD6,38},{12,0xD7,39},{12,0xDA,42},{12,0xDB,43},{13,0x4A,640},{13,0x4B,704},{13,0x4C,768},{13,0x4D,832},{13,0x52,1280},{13,0x53,1344},{13,0x54,1408},{13,0x55,1472},{13,0x5A,1536},{13,0x5B,1600},{13,0x64,1664},{13,0x65,1728},{13,0x6C,512},{13,0x6D,576},{13,0x72,896},{13,0x73,960},{13,0x74,1024},{13,0x75,1088},{13,0x76,1152},{13,0x77,1216}};);func (_dgd *Decoder )detectAndSkipEOL ()error {for {_caa ,_fa :=_dgd ._abd .uncompressGetCode (_dgd ._ceb );if _fa !=nil {return _fa ;};if _caa !=nil &&_caa ._gd ==EOL {_dgd ._abd ._bed +=_caa ._cb ;}else {return nil ;};};};func (_bec *runData )uncompressGetNextCodeLittleEndian ()(int ,error ){_decd :=_bec ._bed -_bec ._edf ;if _decd < 0||_decd > 24{_fgaa :=(_bec ._bed >>3)-_bec ._ec ;if _fgaa >=_bec ._ggc {_fgaa +=_bec ._ec ;if _eed :=_bec .fillBuffer (_fgaa );_eed !=nil {return 0,_eed ;};_fgaa -=_bec ._ec ;};_dbd :=(uint32 (_bec ._edd [_fgaa ]&0xFF)<<16)|(uint32 (_bec ._edd [_fgaa +1]&0xFF)<<8)|(uint32 (_bec ._edd [_fgaa +2]&0xFF));_ffd :=uint32 (_bec ._bed &7);_dbd <<=_ffd ;_bec ._fdf =int (_dbd );}else {_cca :=_bec ._edf &7;_deac :=7-_cca ;if _decd <=_deac {_bec ._fdf <<=uint (_decd );}else {_ggf :=(_bec ._edf >>3)+3-_bec ._ec ;if _ggf >=_bec ._ggc {_ggf +=_bec ._ec ;if _efda :=_bec .fillBuffer (_ggf );_efda !=nil {return 0,_efda ;};_ggf -=_bec ._ec ;};_cca =8-_cca ;for {_bec ._fdf <<=uint (_cca );_bec ._fdf |=int (uint (_bec ._edd [_ggf ])&0xFF);_decd -=_cca ;_ggf ++;_cca =8;if !(_decd >=8){break ;};};_bec ._fdf <<=uint (_decd );};};_bec ._edf =_bec ._bed ;return _bec ._fdf ,nil ;};