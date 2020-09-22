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

package ccittfax ;import (_c "errors";_cc "math";);func (_fg *Encoder )decodeG31D (_eg []byte )([][]byte ,error ){var _cdc [][]byte ;var _dcb int ;for (_dcb /8)< len (_eg ){var _cde bool ;_cde ,_dcb =_gbdc (_eg ,_dcb );if !_cde {if _fg .EndOfLine {return nil ,_cd ;};}else {for _fge :=0;_fge < 5;_fge ++{_cde ,_dcb =_gbdc (_eg ,_dcb );if !_cde {if _fge ==0{break ;};return nil ,_af ;};};if _cde {break ;};};var _aec []byte ;_aec ,_dcb =_fg .decodeRow1D (_eg ,_dcb );if _fg .EncodedByteAlign &&_dcb %8!=0{_dcb +=8-_dcb %8;};_cdc =append (_cdc ,_aec );if _fg .Rows > 0&&!_fg .EndOfBlock &&len (_cdc )>=_fg .Rows {break ;};};return _cdc ,nil ;};var (_ca map[int ]code ;_d map[int ]code ;_g map[int ]code ;_e map[int ]code ;_cce map[int ]code ;_gb map[int ]byte ;_cf =code {Code :1<<4,BitsWritten :12};_de =code {Code :3<<3,BitsWritten :13};_be =code {Code :2<<3,BitsWritten :13};_a =code {Code :1<<12,BitsWritten :4};_bc =code {Code :1<<13,BitsWritten :3};_bcd =code {Code :1<<15,BitsWritten :1};_bd =code {Code :3<<13,BitsWritten :3};_f =code {Code :3<<10,BitsWritten :6};_bce =code {Code :3<<9,BitsWritten :7};_bf =code {Code :2<<13,BitsWritten :3};_fd =code {Code :2<<10,BitsWritten :6};_ed =code {Code :2<<9,BitsWritten :7};);func init (){_ca =make (map[int ]code );_ca [0]=code {Code :13<<8|3<<6,BitsWritten :10};_ca [1]=code {Code :2<<(5+8),BitsWritten :3};_ca [2]=code {Code :3<<(6+8),BitsWritten :2};_ca [3]=code {Code :2<<(6+8),BitsWritten :2};_ca [4]=code {Code :3<<(5+8),BitsWritten :3};_ca [5]=code {Code :3<<(4+8),BitsWritten :4};_ca [6]=code {Code :2<<(4+8),BitsWritten :4};_ca [7]=code {Code :3<<(3+8),BitsWritten :5};_ca [8]=code {Code :5<<(2+8),BitsWritten :6};_ca [9]=code {Code :4<<(2+8),BitsWritten :6};_ca [10]=code {Code :4<<(1+8),BitsWritten :7};_ca [11]=code {Code :5<<(1+8),BitsWritten :7};_ca [12]=code {Code :7<<(1+8),BitsWritten :7};_ca [13]=code {Code :4<<8,BitsWritten :8};_ca [14]=code {Code :7<<8,BitsWritten :8};_ca [15]=code {Code :12<<8,BitsWritten :9};_ca [16]=code {Code :5<<8|3<<6,BitsWritten :10};_ca [17]=code {Code :6<<8,BitsWritten :10};_ca [18]=code {Code :2<<8,BitsWritten :10};_ca [19]=code {Code :12<<8|7<<5,BitsWritten :11};_ca [20]=code {Code :13<<8,BitsWritten :11};_ca [21]=code {Code :13<<8|4<<5,BitsWritten :11};_ca [22]=code {Code :6<<8|7<<5,BitsWritten :11};_ca [23]=code {Code :5<<8,BitsWritten :11};_ca [24]=code {Code :2<<8|7<<5,BitsWritten :11};_ca [25]=code {Code :3<<8,BitsWritten :11};_ca [26]=code {Code :12<<8|10<<4,BitsWritten :12};_ca [27]=code {Code :12<<8|11<<4,BitsWritten :12};_ca [28]=code {Code :12<<8|12<<4,BitsWritten :12};_ca [29]=code {Code :12<<8|13<<4,BitsWritten :12};_ca [30]=code {Code :6<<8|8<<4,BitsWritten :12};_ca [31]=code {Code :6<<8|9<<4,BitsWritten :12};_ca [32]=code {Code :6<<8|10<<4,BitsWritten :12};_ca [33]=code {Code :6<<8|11<<4,BitsWritten :12};_ca [34]=code {Code :13<<8|2<<4,BitsWritten :12};_ca [35]=code {Code :13<<8|3<<4,BitsWritten :12};_ca [36]=code {Code :13<<8|4<<4,BitsWritten :12};_ca [37]=code {Code :13<<8|5<<4,BitsWritten :12};_ca [38]=code {Code :13<<8|6<<4,BitsWritten :12};_ca [39]=code {Code :13<<8|7<<4,BitsWritten :12};_ca [40]=code {Code :6<<8|12<<4,BitsWritten :12};_ca [41]=code {Code :6<<8|13<<4,BitsWritten :12};_ca [42]=code {Code :13<<8|10<<4,BitsWritten :12};_ca [43]=code {Code :13<<8|11<<4,BitsWritten :12};_ca [44]=code {Code :5<<8|4<<4,BitsWritten :12};_ca [45]=code {Code :5<<8|5<<4,BitsWritten :12};_ca [46]=code {Code :5<<8|6<<4,BitsWritten :12};_ca [47]=code {Code :5<<8|7<<4,BitsWritten :12};_ca [48]=code {Code :6<<8|4<<4,BitsWritten :12};_ca [49]=code {Code :6<<8|5<<4,BitsWritten :12};_ca [50]=code {Code :5<<8|2<<4,BitsWritten :12};_ca [51]=code {Code :5<<8|3<<4,BitsWritten :12};_ca [52]=code {Code :2<<8|4<<4,BitsWritten :12};_ca [53]=code {Code :3<<8|7<<4,BitsWritten :12};_ca [54]=code {Code :3<<8|8<<4,BitsWritten :12};_ca [55]=code {Code :2<<8|7<<4,BitsWritten :12};_ca [56]=code {Code :2<<8|8<<4,BitsWritten :12};_ca [57]=code {Code :5<<8|8<<4,BitsWritten :12};_ca [58]=code {Code :5<<8|9<<4,BitsWritten :12};_ca [59]=code {Code :2<<8|11<<4,BitsWritten :12};_ca [60]=code {Code :2<<8|12<<4,BitsWritten :12};_ca [61]=code {Code :5<<8|10<<4,BitsWritten :12};_ca [62]=code {Code :6<<8|6<<4,BitsWritten :12};_ca [63]=code {Code :6<<8|7<<4,BitsWritten :12};_d =make (map[int ]code );_d [0]=code {Code :53<<8,BitsWritten :8};_d [1]=code {Code :7<<(2+8),BitsWritten :6};_d [2]=code {Code :7<<(4+8),BitsWritten :4};_d [3]=code {Code :8<<(4+8),BitsWritten :4};_d [4]=code {Code :11<<(4+8),BitsWritten :4};_d [5]=code {Code :12<<(4+8),BitsWritten :4};_d [6]=code {Code :14<<(4+8),BitsWritten :4};_d [7]=code {Code :15<<(4+8),BitsWritten :4};_d [8]=code {Code :19<<(3+8),BitsWritten :5};_d [9]=code {Code :20<<(3+8),BitsWritten :5};_d [10]=code {Code :7<<(3+8),BitsWritten :5};_d [11]=code {Code :8<<(3+8),BitsWritten :5};_d [12]=code {Code :8<<(2+8),BitsWritten :6};_d [13]=code {Code :3<<(2+8),BitsWritten :6};_d [14]=code {Code :52<<(2+8),BitsWritten :6};_d [15]=code {Code :53<<(2+8),BitsWritten :6};_d [16]=code {Code :42<<(2+8),BitsWritten :6};_d [17]=code {Code :43<<(2+8),BitsWritten :6};_d [18]=code {Code :39<<(1+8),BitsWritten :7};_d [19]=code {Code :12<<(1+8),BitsWritten :7};_d [20]=code {Code :8<<(1+8),BitsWritten :7};_d [21]=code {Code :23<<(1+8),BitsWritten :7};_d [22]=code {Code :3<<(1+8),BitsWritten :7};_d [23]=code {Code :4<<(1+8),BitsWritten :7};_d [24]=code {Code :40<<(1+8),BitsWritten :7};_d [25]=code {Code :43<<(1+8),BitsWritten :7};_d [26]=code {Code :19<<(1+8),BitsWritten :7};_d [27]=code {Code :36<<(1+8),BitsWritten :7};_d [28]=code {Code :24<<(1+8),BitsWritten :7};_d [29]=code {Code :2<<8,BitsWritten :8};_d [30]=code {Code :3<<8,BitsWritten :8};_d [31]=code {Code :26<<8,BitsWritten :8};_d [32]=code {Code :27<<8,BitsWritten :8};_d [33]=code {Code :18<<8,BitsWritten :8};_d [34]=code {Code :19<<8,BitsWritten :8};_d [35]=code {Code :20<<8,BitsWritten :8};_d [36]=code {Code :21<<8,BitsWritten :8};_d [37]=code {Code :22<<8,BitsWritten :8};_d [38]=code {Code :23<<8,BitsWritten :8};_d [39]=code {Code :40<<8,BitsWritten :8};_d [40]=code {Code :41<<8,BitsWritten :8};_d [41]=code {Code :42<<8,BitsWritten :8};_d [42]=code {Code :43<<8,BitsWritten :8};_d [43]=code {Code :44<<8,BitsWritten :8};_d [44]=code {Code :45<<8,BitsWritten :8};_d [45]=code {Code :4<<8,BitsWritten :8};_d [46]=code {Code :5<<8,BitsWritten :8};_d [47]=code {Code :10<<8,BitsWritten :8};_d [48]=code {Code :11<<8,BitsWritten :8};_d [49]=code {Code :82<<8,BitsWritten :8};_d [50]=code {Code :83<<8,BitsWritten :8};_d [51]=code {Code :84<<8,BitsWritten :8};_d [52]=code {Code :85<<8,BitsWritten :8};_d [53]=code {Code :36<<8,BitsWritten :8};_d [54]=code {Code :37<<8,BitsWritten :8};_d [55]=code {Code :88<<8,BitsWritten :8};_d [56]=code {Code :89<<8,BitsWritten :8};_d [57]=code {Code :90<<8,BitsWritten :8};_d [58]=code {Code :91<<8,BitsWritten :8};_d [59]=code {Code :74<<8,BitsWritten :8};_d [60]=code {Code :75<<8,BitsWritten :8};_d [61]=code {Code :50<<8,BitsWritten :8};_d [62]=code {Code :51<<8,BitsWritten :8};_d [63]=code {Code :52<<8,BitsWritten :8};_g =make (map[int ]code );_g [64]=code {Code :3<<8|3<<6,BitsWritten :10};_g [128]=code {Code :12<<8|8<<4,BitsWritten :12};_g [192]=code {Code :12<<8|9<<4,BitsWritten :12};_g [256]=code {Code :5<<8|11<<4,BitsWritten :12};_g [320]=code {Code :3<<8|3<<4,BitsWritten :12};_g [384]=code {Code :3<<8|4<<4,BitsWritten :12};_g [448]=code {Code :3<<8|5<<4,BitsWritten :12};_g [512]=code {Code :3<<8|12<<3,BitsWritten :13};_g [576]=code {Code :3<<8|13<<3,BitsWritten :13};_g [640]=code {Code :2<<8|10<<3,BitsWritten :13};_g [704]=code {Code :2<<8|11<<3,BitsWritten :13};_g [768]=code {Code :2<<8|12<<3,BitsWritten :13};_g [832]=code {Code :2<<8|13<<3,BitsWritten :13};_g [896]=code {Code :3<<8|18<<3,BitsWritten :13};_g [960]=code {Code :3<<8|19<<3,BitsWritten :13};_g [1024]=code {Code :3<<8|20<<3,BitsWritten :13};_g [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_g [1152]=code {Code :3<<8|22<<3,BitsWritten :13};_g [1216]=code {Code :119<<3,BitsWritten :13};_g [1280]=code {Code :2<<8|18<<3,BitsWritten :13};_g [1344]=code {Code :2<<8|19<<3,BitsWritten :13};_g [1408]=code {Code :2<<8|20<<3,BitsWritten :13};_g [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_g [1536]=code {Code :2<<8|26<<3,BitsWritten :13};_g [1600]=code {Code :2<<8|27<<3,BitsWritten :13};_g [1664]=code {Code :3<<8|4<<3,BitsWritten :13};_g [1728]=code {Code :3<<8|5<<3,BitsWritten :13};_e =make (map[int ]code );_e [64]=code {Code :27<<(3+8),BitsWritten :5};_e [128]=code {Code :18<<(3+8),BitsWritten :5};_e [192]=code {Code :23<<(2+8),BitsWritten :6};_e [256]=code {Code :55<<(1+8),BitsWritten :7};_e [320]=code {Code :54<<8,BitsWritten :8};_e [384]=code {Code :55<<8,BitsWritten :8};_e [448]=code {Code :100<<8,BitsWritten :8};_e [512]=code {Code :101<<8,BitsWritten :8};_e [576]=code {Code :104<<8,BitsWritten :8};_e [640]=code {Code :103<<8,BitsWritten :8};_e [704]=code {Code :102<<8,BitsWritten :9};_e [768]=code {Code :102<<8|1<<7,BitsWritten :9};_e [832]=code {Code :105<<8,BitsWritten :9};_e [896]=code {Code :105<<8|1<<7,BitsWritten :9};_e [960]=code {Code :106<<8,BitsWritten :9};_e [1024]=code {Code :106<<8|1<<7,BitsWritten :9};_e [1088]=code {Code :107<<8,BitsWritten :9};_e [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_e [1216]=code {Code :108<<8,BitsWritten :9};_e [1280]=code {Code :108<<8|1<<7,BitsWritten :9};_e [1344]=code {Code :109<<8,BitsWritten :9};_e [1408]=code {Code :109<<8|1<<7,BitsWritten :9};_e [1472]=code {Code :76<<8,BitsWritten :9};_e [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_e [1600]=code {Code :77<<8,BitsWritten :9};_e [1664]=code {Code :24<<(2+8),BitsWritten :6};_e [1728]=code {Code :77<<8|1<<7,BitsWritten :9};_cce =make (map[int ]code );_cce [1792]=code {Code :1<<8,BitsWritten :11};_cce [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_cce [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_cce [1984]=code {Code :1<<8|2<<4,BitsWritten :12};_cce [2048]=code {Code :1<<8|3<<4,BitsWritten :12};_cce [2112]=code {Code :1<<8|4<<4,BitsWritten :12};_cce [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_cce [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_cce [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_cce [2368]=code {Code :1<<8|12<<4,BitsWritten :12};_cce [2432]=code {Code :1<<8|13<<4,BitsWritten :12};_cce [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_cce [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_gb =make (map[int ]byte );_gb [0]=0xFF;_gb [1]=0xFE;_gb [2]=0xFC;_gb [3]=0xF8;_gb [4]=0xF0;_gb [5]=0xE0;_gb [6]=0xC0;_gb [7]=0x80;_gb [8]=0x00;};func (_dfda *Encoder )encodeG32D (_aadb [][]byte )[]byte {var _ecbe []byte ;var _gcc int ;for _dfea :=0;_dfea < len (_aadb );_dfea +=_dfda .K {if _dfda .Rows > 0&&!_dfda .EndOfBlock &&_dfea ==_dfda .Rows {break ;};_gfbd ,_ffd :=_abcg (_aadb [_dfea ],_gcc ,_de );_ecbe =_dfda .appendEncodedRow (_ecbe ,_gfbd ,_gcc );if _dfda .EncodedByteAlign {_ffd =0;};_gcc =_ffd ;for _bda :=_dfea +1;_bda < (_dfea +_dfda .K )&&_bda < len (_aadb );_bda ++{if _dfda .Rows > 0&&!_dfda .EndOfBlock &&_bda ==_dfda .Rows {break ;};_cafa ,_ccf :=_gaeb (nil ,_gcc ,_be );var _bga ,_bfbg ,_aae int ;_ebcf :=-1;for _ebcf < len (_aadb [_bda ]){_bga =_cec (_aadb [_bda ],_ebcf );_bfbg =_cfc (_aadb [_bda ],_aadb [_bda -1],_ebcf );_aae =_cec (_aadb [_bda -1],_bfbg );if _aae < _bga {_cafa ,_ccf =_agd (_cafa ,_ccf );_ebcf =_aae ;}else {if _cc .Abs (float64 (_bfbg -_bga ))> 3{_cafa ,_ccf ,_ebcf =_ccfe (_aadb [_bda ],_cafa ,_ccf ,_ebcf ,_bga );}else {_cafa ,_ccf =_fddb (_cafa ,_ccf ,_bga ,_bfbg );_ebcf =_bga ;};};};_ecbe =_dfda .appendEncodedRow (_ecbe ,_cafa ,_gcc );if _dfda .EncodedByteAlign {_ccf =0;};_gcc =_ccf %8;};};if _dfda .EndOfBlock {_gea ,_ :=_ffg (_gcc );_ecbe =_dfda .appendEncodedRow (_ecbe ,_gea ,_gcc );};return _ecbe ;};func (_aaa *Encoder )decodeG4 (_fee []byte )([][]byte ,error ){_ccg :=make ([]byte ,_aaa .Columns );for _aba :=range _ccg {_ccg [_aba ]=_ccgf ;};_faa :=make ([][]byte ,1);_faa [0]=_ccg ;var (_db bool ;_edc error ;_cgg int ;);for (_cgg /8)< len (_fee ){_db ,_cgg ,_edc =_gfb (_fee ,_cgg );if _edc !=nil {return nil ,_edc ;};if _db {break ;};var (_ad code ;_eb bool ;);_aaad :=true ;var _dee []byte ;_dbf :=-1;for _dbf < _aaa .Columns {_ad ,_cgg ,_eb =_cdf (_fee ,_cgg );if !_eb {return nil ,_aa ;};switch _ad {case _a :_dee ,_dbf =_aeb (_faa ,_dee ,_aaad ,_dbf );case _bc :_dee ,_cgg ,_dbf ,_edc =_abf (_fee ,_dee ,_cgg ,_aaad ,_dbf );if _edc !=nil {return nil ,_edc ;};case _bcd :_dee ,_dbf =_adc (_faa ,_dee ,_aaad ,_dbf ,0);_aaad =!_aaad ;case _bd :_dee ,_dbf =_adc (_faa ,_dee ,_aaad ,_dbf ,1);_aaad =!_aaad ;case _f :_dee ,_dbf =_adc (_faa ,_dee ,_aaad ,_dbf ,2);_aaad =!_aaad ;case _bce :_dee ,_dbf =_adc (_faa ,_dee ,_aaad ,_dbf ,3);_aaad =!_aaad ;case _bf :_dee ,_dbf =_adc (_faa ,_dee ,_aaad ,_dbf ,-1);_aaad =!_aaad ;case _fd :_dee ,_dbf =_adc (_faa ,_dee ,_aaad ,_dbf ,-2);_aaad =!_aaad ;case _ed :_dee ,_dbf =_adc (_faa ,_dee ,_aaad ,_dbf ,-3);_aaad =!_aaad ;};if len (_dee )>=_aaa .Columns {break ;};};if _aaa .EncodedByteAlign &&_cgg %8!=0{_cgg +=8-_cgg %8;};_faa =append (_faa ,_dee );if _aaa .Rows > 0&&!_aaa .EndOfBlock &&len (_faa )>=(_aaa .Rows +1){break ;};};_faa =_faa [1:];return _faa ,nil ;};func init (){for _dg ,_ab :=range _d {_caeg (_cac ,_ab ,0,_dg );};for _bed ,_dc :=range _e {_caeg (_cac ,_dc ,0,_bed );};for _df ,_ga :=range _ca {_caeg (_cfd ,_ga ,0,_df );};for _edd ,_fc :=range _g {_caeg (_cfd ,_fc ,0,_edd );};for _cg ,_ccd :=range _cce {_caeg (_cac ,_ccd ,0,_cg );_caeg (_cfd ,_ccd ,0,_cg );};_caeg (_ae ,_a ,0,0);_caeg (_ae ,_bc ,0,0);_caeg (_ae ,_bcd ,0,0);_caeg (_ae ,_bd ,0,0);_caeg (_ae ,_f ,0,0);_caeg (_ae ,_bce ,0,0);_caeg (_ae ,_bf ,0,0);_caeg (_ae ,_fd ,0,0);_caeg (_ae ,_ed ,0,0);};func _ged (_dbff []byte ,_caeb int ,_adb code )(bool ,int ){_dge :=_caeb ;var (_bfb uint16 ;_deda int ;);_bfb ,_deda ,_caeb =_bac (_dbff ,_caeb );if _deda > 3{return false ,_dge ;};_bfb >>=uint (3-_deda );_bfb <<=3;if _bfb !=_adb .Code {return false ,_dge ;}else {return true ,_caeb -3+_deda ;};};func _gbdc (_bedg []byte ,_fgb int )(bool ,int ){_gdc :=_fgb ;var (_gbdb uint16 ;_fbg int ;);_gbdb ,_fbg ,_fgb =_bac (_bedg ,_fgb );if _fbg > 4{return false ,_gdc ;};_gbdb >>=uint (4-_fbg );_gbdb <<=4;if _gbdb !=_cf .Code {return false ,_gdc ;}else {return true ,_fgb -4+_fbg ;};};func _fddb (_daa []byte ,_edcd ,_fdcc ,_acg int )([]byte ,int ){_agac :=_beed (_fdcc ,_acg );_daa ,_edcd =_gaeb (_daa ,_edcd ,_agac );return _daa ,_edcd ;};type code struct{Code uint16 ;BitsWritten int ;};func _caeg (_gbdg *decodingTreeNode ,_dbfe code ,_ccba int ,_agg int ){_abd :=_cae (_dbfe .Code ,_ccba );_ccba ++;if _abd ==1{if _gbdg .Right ==nil {_gbdg .Right =&decodingTreeNode {Val :_abd };};if _ccba ==_dbfe .BitsWritten {_gbdg .Right .RunLen =&_agg ;_gbdg .Right .Code =&_dbfe ;}else {_caeg (_gbdg .Right ,_dbfe ,_ccba ,_agg );};}else {if _gbdg .Left ==nil {_gbdg .Left =&decodingTreeNode {Val :_abd };};if _ccba ==_dbfe .BitsWritten {_gbdg .Left .RunLen =&_agg ;_gbdg .Left .Code =&_dbfe ;}else {_caeg (_gbdg .Left ,_dbfe ,_ccba ,_agg );};};};func _beed (_fcf ,_add int )code {var _afbb code ;switch _add -_fcf {case -1:_afbb =_bd ;case -2:_afbb =_f ;case -3:_afbb =_bce ;case 0:_afbb =_bcd ;case 1:_afbb =_bf ;case 2:_afbb =_fd ;case 3:_afbb =_ed ;};return _afbb ;};func _abfe (_dgb uint16 ,_ee int )(code ,bool ){_ ,_eac :=_fcg (_ae ,_dgb ,_ee );if _eac ==nil {return code {},false ;};return *_eac ,true ;};func _bac (_bdb []byte ,_ccb int )(uint16 ,int ,int ){_gee :=_ccb ;_gabe :=_ccb /8;_ccb %=8;if _gabe >=len (_bdb ){return 0,16,_gee ;};_dbb :=byte (0xFF>>uint (_ccb ));_dce :=uint16 ((_bdb [_gabe ]&_dbb )<<uint (_ccb ))<<8;_cffg :=8-_ccb ;_gabe ++;_ccb =0;if _gabe >=len (_bdb ){return _dce >>(16-uint (_cffg )),16-_cffg ,_gee +_cffg ;};_dce |=uint16 (_bdb [_gabe ])<<(8-uint (_cffg ));_cffg +=8;_gabe ++;_ccb =0;if _gabe >=len (_bdb ){return _dce >>(16-uint (_cffg )),16-_cffg ,_gee +_cffg ;};if _cffg ==16{return _dce ,0,_gee +_cffg ;};_gfc :=16-_cffg ;_dce |=uint16 (_bdb [_gabe ]>>(8-uint (_gfc )));return _dce ,0,_gee +16;};func _aga (_fdf ,_eggd []byte ,_eef int ,_ecea bool )int {_fbdb :=_cec (_eggd ,_eef );if _fbdb < len (_eggd )&&(_eef ==-1&&_eggd [_fbdb ]==_ccgf ||_eef >=0&&_eef < len (_fdf )&&_fdf [_eef ]==_eggd [_fbdb ]||_eef >=len (_fdf )&&_ecea &&_eggd [_fbdb ]==_ccgf ||_eef >=len (_fdf )&&!_ecea &&_eggd [_fbdb ]==_dag ){_fbdb =_cec (_eggd ,_fbdb );};return _fbdb ;};func _dedf (_gab []byte ,_bgb int )(bool ,int ,error ){_ccea :=_bgb ;var _cbe =false ;for _edg :=0;_edg < 6;_edg ++{_cbe ,_bgb =_ece (_gab ,_bgb );if !_cbe {if _edg > 1{return false ,_ccea ,_ge ;}else {_bgb =_ccea ;break ;};};};return _cbe ,_bgb ,nil ;};func _cfc (_bcga ,_afb []byte ,_bdfe int )int {_abb :=_cec (_afb ,_bdfe );if _abb < len (_afb )&&(_bdfe ==-1&&_afb [_abb ]==_ccgf ||_bdfe >=0&&_bdfe < len (_bcga )&&_bcga [_bdfe ]==_afb [_abb ]||_bdfe >=len (_bcga )&&_bcga [_bdfe -1]!=_afb [_abb ]){_abb =_cec (_afb ,_abb );};return _abb ;};var (_cac =&decodingTreeNode {Val :255};_cfd =&decodingTreeNode {Val :255};_ae =&decodingTreeNode {Val :255};);func _fcg (_ecb *decodingTreeNode ,_fff uint16 ,_gfd int )(*int ,*code ){if _ecb ==nil {return nil ,nil ;};if _gfd ==16{return _ecb .RunLen ,_ecb .Code ;};_cdae :=_cae (_fff ,_gfd );_gfd ++;var _dgd *int ;var _ccc *code ;if _cdae ==1{_dgd ,_ccc =_fcg (_ecb .Right ,_fff ,_gfd );}else {_dgd ,_ccc =_fcg (_ecb .Left ,_fff ,_gfd );};if _dgd ==nil {_dgd =_ecb .RunLen ;_ccc =_ecb .Code ;};return _dgd ,_ccc ;};var (_ccgf byte =1;_dag byte =0;);func _adf (_adbd []byte ,_dcef int )(bool ,int ){return _ged (_adbd ,_dcef ,_be )};func _gfb (_fec []byte ,_dgg int )(bool ,int ,error ){_dggd :=_dgg ;var _dfe bool ;_dfe ,_dgg =_gbdc (_fec ,_dgg );if _dfe {_dfe ,_dgg =_gbdc (_fec ,_dgg );if _dfe {return true ,_dgg ,nil ;}else {return false ,_dggd ,_ded ;};};return false ,_dggd ,nil ;};func _abf (_gfa ,_dcf []byte ,_gg int ,_fdbe bool ,_fb int )([]byte ,int ,int ,error ){_cbf :=_gg ;var _bdf error ;_dcf ,_gg ,_bdf =_dfd (_gfa ,_dcf ,_gg ,_fdbe );if _bdf !=nil {return _dcf ,_cbf ,_fb ,_bdf ;};_fdbe =!_fdbe ;_dcf ,_gg ,_bdf =_dfd (_gfa ,_dcf ,_gg ,_fdbe );if _bdf !=nil {return _dcf ,_cbf ,_fb ,_bdf ;};_fb =len (_dcf );return _dcf ,_gg ,_fb ,nil ;};func _cae (_dbg uint16 ,_cabf int )byte {if _cabf < 8{_dbg >>=8;};_cabf %=8;_gga :=byte (0x01<<(7-uint (_cabf )));return (byte (_dbg )&_gga )>>(7-uint (_cabf ));};func (_fbfa *Encoder )Encode (pixels [][]byte )[]byte {if _fbfa .BlackIs1 {_ccgf =0;_dag =1;}else {_ccgf =1;_dag =0;};if _fbfa .K ==0{return _fbfa .encodeG31D (pixels );};if _fbfa .K > 0{return _fbfa .encodeG32D (pixels );};if _fbfa .K < 0{return _fbfa .encodeG4 (pixels );};return nil ;};func _fgd (_cbea []byte ,_ggcb int ,_gae int ,_fea bool )([]byte ,int ){var (_ecg code ;_feb bool ;);for !_feb {_ecg ,_gae ,_feb =_bgee (_gae ,_fea );_cbea ,_ggcb =_gaeb (_cbea ,_ggcb ,_ecg );};return _cbea ,_ggcb ;};func (_gead *Encoder )appendEncodedRow (_geg ,_bad []byte ,_dca int )[]byte {if len (_geg )> 0&&_dca !=0&&!_gead .EncodedByteAlign {_geg [len (_geg )-1]=_geg [len (_geg )-1]|_bad [0];_geg =append (_geg ,_bad [1:]...);}else {_geg =append (_geg ,_bad ...);};return _geg ;};func _gad (_bfa int )([]byte ,int ){var _acb []byte ;for _egg :=0;_egg < 2;_egg ++{_acb ,_bfa =_gaeb (_acb ,_bfa ,_cf );};return _acb ,_bfa %8;};func _gce (_cceg int )([]byte ,int ){var _eba []byte ;for _gbe :=0;_gbe < 6;_gbe ++{_eba ,_cceg =_gaeb (_eba ,_cceg ,_cf );};return _eba ,_cceg %8;};func _cacb (_ffc []byte ,_gcd bool ,_cafag int )(int ,int ){_cad :=0;for _cafag < len (_ffc ){if _gcd {if _ffc [_cafag ]!=_ccgf {break ;};}else {if _ffc [_cafag ]!=_dag {break ;};};_cad ++;_cafag ++;};return _cad ,_cafag ;};func _dfd (_gag ,_bcb []byte ,_beg int ,_ade bool )([]byte ,int ,error ){_aef :=_beg ;var _ccdb int ;for _ccdb ,_beg =_baa (_gag ,_beg ,_ade );_ccdb !=-1;_ccdb ,_beg =_baa (_gag ,_beg ,_ade ){_bcb =_da (_bcb ,_ade ,_ccdb );if _ccdb < 64{break ;};};if _ccdb ==-1{return _bcb ,_aef ,_cff ;};return _bcb ,_beg ,nil ;};type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;EndOfBlock bool ;BlackIs1 bool ;DamagedRowsBeforeError int ;};func _gaa (_ecd uint16 ,_fecf int ,_fbf bool )(int ,code ){var _bbg *int ;var _ff *code ;if _fbf {_bbg ,_ff =_fcg (_cac ,_ecd ,_fecf );}else {_bbg ,_ff =_fcg (_cfd ,_ecd ,_fecf );};if _bbg ==nil {return -1,code {};};return *_bbg ,*_ff ;};func _ece (_cda []byte ,_afca int )(bool ,int ){return _ged (_cda ,_afca ,_de )};func _bgee (_afac int ,_bdc bool )(code ,int ,bool ){if _afac < 64{if _bdc {return _d [_afac ],0,true ;}else {return _ca [_afac ],0,true ;};}else {_agf :=_afac /64;if _agf > 40{return _cce [2560],_afac -2560,false ;};if _agf > 27{return _cce [_agf *64],_afac -_agf *64,false ;};if _bdc {return _e [_agf *64],_afac -_agf *64,false ;}else {return _g [_agf *64],_afac -_agf *64,false ;};};};func (_efb *Encoder )encodeG4 (_cee [][]byte )[]byte {_fde :=make ([][]byte ,len (_cee ));copy (_fde ,_cee );_fde =_fecg (_fde );var _dfec []byte ;var _abc int ;for _efe :=1;_efe < len (_fde );_efe ++{if _efb .Rows > 0&&!_efb .EndOfBlock &&_efe ==(_efb .Rows +1){break ;};var _eec []byte ;var _fab ,_cbc ,_dbfdd int ;_bca :=_abc ;_cge :=-1;for _cge < len (_fde [_efe ]){_fab =_cec (_fde [_efe ],_cge );_cbc =_cfc (_fde [_efe ],_fde [_efe -1],_cge );_dbfdd =_cec (_fde [_efe -1],_cbc );if _dbfdd < _fab {_eec ,_bca =_gaeb (_eec ,_bca ,_a );_cge =_dbfdd ;}else {if _cc .Abs (float64 (_cbc -_fab ))> 3{_eec ,_bca ,_cge =_ccfe (_fde [_efe ],_eec ,_bca ,_cge ,_fab );}else {_eec ,_bca =_fddb (_eec ,_bca ,_fab ,_cbc );_cge =_fab ;};};};_dfec =_efb .appendEncodedRow (_dfec ,_eec ,_abc );if _efb .EncodedByteAlign {_bca =0;};_abc =_bca %8;};if _efb .EndOfBlock {_cca ,_ :=_gad (_abc );_dfec =_efb .appendEncodedRow (_dfec ,_cca ,_abc );};return _dfec ;};func _da (_gdd []byte ,_cffc bool ,_ada int )[]byte {if _ada < 0{return _gdd ;};_dgc :=make ([]byte ,_ada );if _cffc {for _aacea :=0;_aacea < len (_dgc );_aacea ++{_dgc [_aacea ]=_ccgf ;};}else {for _cga :=0;_cga < len (_dgc );_cga ++{_dgc [_cga ]=_dag ;};};_gdd =append (_gdd ,_dgc ...);return _gdd ;};func _abcg (_cfb []byte ,_gef int ,_dfc code )([]byte ,int ){_aea :=true ;var _adbg []byte ;_adbg ,_gef =_gaeb (nil ,_gef ,_dfc );_adcf :=0;var _edcg int ;for _adcf < len (_cfb ){_edcg ,_adcf =_cacb (_cfb ,_aea ,_adcf );_adbg ,_gef =_fgd (_adbg ,_gef ,_edcg ,_aea );_aea =!_aea ;};return _adbg ,_gef %8;};func (_fdd *Encoder )decodeRow1D (_fbb []byte ,_fac int )([]byte ,int ){var _bfd []byte ;_gbdf :=true ;var _bee int ;_bee ,_fac =_baa (_fbb ,_fac ,_gbdf );for _bee !=-1{_bfd =_da (_bfd ,_gbdf ,_bee );if _bee < 64{if len (_bfd )>=_fdd .Columns {break ;};_gbdf =!_gbdf ;};_bee ,_fac =_baa (_fbb ,_fac ,_gbdf );};return _bfd ,_fac ;};func (_adg *Encoder )encodeG31D (_aad [][]byte )[]byte {var _cdee []byte ;_fcd :=0;for _caf :=range _aad {if _adg .Rows > 0&&!_adg .EndOfBlock &&_caf ==_adg .Rows {break ;};_dcee ,_beag :=_abcg (_aad [_caf ],_fcd ,_cf );_cdee =_adg .appendEncodedRow (_cdee ,_dcee ,_fcd );if _adg .EncodedByteAlign {_beag =0;};_fcd =_beag ;};if _adg .EndOfBlock {_dbfd ,_ :=_gce (_fcd );_cdee =_adg .appendEncodedRow (_cdee ,_dbfd ,_fcd );};return _cdee ;};func _aeb (_bbf [][]byte ,_cab []byte ,_ac bool ,_cb int )([]byte ,int ){_aace :=_aga (_cab ,_bbf [len (_bbf )-1],_cb ,_ac );_bcg :=_cec (_bbf [len (_bbf )-1],_aace );if _cb ==-1{_cab =_da (_cab ,_ac ,_bcg -_cb -1);}else {_cab =_da (_cab ,_ac ,_bcg -_cb );};_cb =_bcg ;return _cab ,_cb ;};func _fecg (_gfde [][]byte )[][]byte {_ecab :=make ([]byte ,len (_gfde [0]));for _geadg :=range _ecab {_ecab [_geadg ]=_ccgf ;};_gfde =append (_gfde ,[]byte {});for _dfdf :=len (_gfde )-1;_dfdf > 0;_dfdf --{_gfde [_dfdf ]=_gfde [_dfdf -1];};_gfde [0]=_ecab ;return _gfde ;};type decodingTreeNode struct{Val byte ;RunLen *int ;Code *code ;Left *decodingTreeNode ;Right *decodingTreeNode ;};func _agd (_gcea []byte ,_dcbd int )([]byte ,int ){return _gaeb (_gcea ,_dcbd ,_a )};func _ffg (_eecf int )([]byte ,int ){var _cfg []byte ;for _bge :=0;_bge < 6;_bge ++{_cfg ,_eecf =_gaeb (_cfg ,_eecf ,_de );};return _cfg ,_eecf %8;};func _cec (_fgde []byte ,_efa int )int {if _efa >=len (_fgde ){return _efa ;};if _efa < -1{_efa =-1;};var _fbc byte ;if _efa > -1{_fbc =_fgde [_efa ];}else {_fbc =_ccgf ;};_fdc :=_efa +1;for _fdc < len (_fgde ){if _fgde [_fdc ]!=_fbc {break ;};_fdc ++;};return _fdc ;};func _cdf (_beb []byte ,_gabf int )(code ,int ,bool ){var (_egc uint16 ;_ggc int ;_gbd int ;);_gbd =_gabf ;_egc ,_ggc ,_gabf =_bac (_beb ,_gabf );_faf ,_cgb :=_abfe (_egc ,_ggc );if !_cgb {return code {},_gbd ,false ;};return _faf ,_gbd +_faf .BitsWritten ,true ;};func _adc (_ebc [][]byte ,_ebg []byte ,_ba bool ,_bef ,_aac int )([]byte ,int ){_ag :=_aga (_ebg ,_ebc [len (_ebc )-1],_bef ,_ba );_ce :=_ag +_aac ;if _bef ==-1{_ebg =_da (_ebg ,_ba ,_ce -_bef -1);}else {_ebg =_da (_ebg ,_ba ,_ce -_bef );};_bef =_ce ;return _ebg ,_bef ;};func _baa (_ebe []byte ,_ecc int ,_afa bool )(int ,int ){var (_ef uint16 ;_fef int ;_bgd int ;);_bgd =_ecc ;_ef ,_fef ,_ecc =_bac (_ebe ,_ecc );_acd ,_dbd :=_gaa (_ef ,_fef ,_afa );if _acd ==-1{return -1,_bgd ;};return _acd ,_bgd +_dbd .BitsWritten ;};var (_ded =_c .New ("\u0045\u004f\u0046\u0042 c\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074e\u0064");_ge =_c .New ("R\u0054\u0043\u0020\u0063od\u0065 \u0069\u0073\u0020\u0063\u006fr\u0072\u0075\u0070\u0074\u0065\u0064");_cff =_c .New ("\u0077\u0072o\u006e\u0067\u0020\u0063\u006f\u0064\u0065\u0020\u0069\u006e\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020mo\u0064\u0065");_cd =_c .New ("\u006e\u006f\u0020\u0045\u004f\u004c\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0074\u0068\u0065\u0020\u0045\u006e\u0064O\u0066\u004c\u0069\u006e\u0065\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0065");_af =_c .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u004f\u004c");_aa =_c .New ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0032\u0044\u0020\u0063\u006f\u0064\u0065"););func (_bg *Encoder )Decode (encoded []byte )([][]byte ,error ){if _bg .BlackIs1 {_ccgf =0;_dag =1;}else {_ccgf =1;_dag =0;};if _bg .K ==0{return _bg .decodeG31D (encoded );};if _bg .K > 0{return _bg .decodeG32D (encoded );};if _bg .K < 4{return _bg .decodeG4 (encoded );};return nil ,nil ;};func _gaeb (_eaf []byte ,_gaf int ,_dcg code )([]byte ,int ){_fgdf :=0;for _fgdf < _dcg .BitsWritten {_dbgc :=_gaf /8;_gdcf :=_gaf %8;if _dbgc >=len (_eaf ){_eaf =append (_eaf ,0);};_fbd :=8-_gdcf ;_gaba :=_dcg .BitsWritten -_fgdf ;if _fbd > _gaba {_fbd =_gaba ;};if _fgdf < 8{_eaf [_dbgc ]=_eaf [_dbgc ]|byte (_dcg .Code >>uint (8+_gdcf -_fgdf ))&_gb [8-_fbd -_gdcf ];}else {_eaf [_dbgc ]=_eaf [_dbgc ]|(byte (_dcg .Code <<uint (_fgdf -8))&_gb [8-_fbd ])>>uint (_gdcf );};_gaf +=_fbd ;_fgdf +=_fbd ;};return _eaf ,_gaf ;};func _ccfe (_ebb ,_dd []byte ,_bfac ,_fbgd ,_cccc int )([]byte ,int ,int ){_adag :=_cec (_ebb ,_cccc );_adagb :=_fbgd >=0&&_ebb [_fbgd ]==_ccgf ||_fbgd ==-1;_dd ,_bfac =_gaeb (_dd ,_bfac ,_bc );var _eece int ;if _fbgd > -1{_eece =_cccc -_fbgd ;}else {_eece =_cccc -_fbgd -1;};_dd ,_bfac =_fgd (_dd ,_bfac ,_eece ,_adagb );_adagb =!_adagb ;_gff :=_adag -_cccc ;_dd ,_bfac =_fgd (_dd ,_bfac ,_gff ,_adagb );_fbgd =_adag ;return _dd ,_bfac ,_fbgd ;};func (_ege *Encoder )decodeG32D (_gf []byte )([][]byte ,error ){var (_bgg [][]byte ;_gd int ;_ec error ;);_fdb :for (_gd /8)< len (_gf ){var _fe bool ;_fe ,_gd ,_ec =_dedf (_gf ,_gd );if _ec !=nil {return nil ,_ec ;};if _fe {break ;};_fe ,_gd =_ece (_gf ,_gd );if !_fe {if _ege .EndOfLine {return nil ,_cd ;};};var _fa []byte ;_fa ,_gd =_ege .decodeRow1D (_gf ,_gd );if _ege .EncodedByteAlign &&_gd %8!=0{_gd +=8-_gd %8;};if _fa !=nil {_bgg =append (_bgg ,_fa );};if _ege .Rows > 0&&!_ege .EndOfBlock &&len (_bgg )>=_ege .Rows {break ;};for _bb :=1;_bb < _ege .K &&(_gd /8)< len (_gf );_bb ++{_fe ,_gd =_adf (_gf ,_gd );if !_fe {_fe ,_gd ,_ec =_dedf (_gf ,_gd );if _ec !=nil {return nil ,_ec ;};if _fe {break _fdb ;}else {if _ege .EndOfLine {return nil ,_cd ;};};};var (_afc code ;_gdg bool ;);_ega :=true ;var _ea []byte ;_eca :=-1;for _afc ,_gd ,_gdg =_cdf (_gf ,_gd );_gdg ;_afc ,_gd ,_gdg =_cdf (_gf ,_gd ){switch _afc {case _a :_ea ,_eca =_aeb (_bgg ,_ea ,_ega ,_eca );case _bc :_ea ,_gd ,_eca ,_ec =_abf (_gf ,_ea ,_gd ,_ega ,_eca );if _ec !=nil {return nil ,_ec ;};case _bcd :_ea ,_eca =_adc (_bgg ,_ea ,_ega ,_eca ,0);_ega =!_ega ;case _bd :_ea ,_eca =_adc (_bgg ,_ea ,_ega ,_eca ,1);_ega =!_ega ;case _f :_ea ,_eca =_adc (_bgg ,_ea ,_ega ,_eca ,2);_ega =!_ega ;case _bce :_ea ,_eca =_adc (_bgg ,_ea ,_ega ,_eca ,3);_ega =!_ega ;case _bf :_ea ,_eca =_adc (_bgg ,_ea ,_ega ,_eca ,-1);_ega =!_ega ;case _fd :_ea ,_eca =_adc (_bgg ,_ea ,_ega ,_eca ,-2);_ega =!_ega ;case _ed :_ea ,_eca =_adc (_bgg ,_ea ,_ega ,_eca ,-3);_ega =!_ega ;};if len (_ea )>=_ege .Columns {break ;};};if _ege .EncodedByteAlign &&_gd %8!=0{_gd +=8-_gd %8;};if _ea !=nil {_bgg =append (_bgg ,_ea );};if _ege .Rows > 0&&!_ege .EndOfBlock &&len (_bgg )>=_ege .Rows {break _fdb ;};};};return _bgg ,nil ;};