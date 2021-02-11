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

package sampling ;import (_a "github.com/unidoc/unipdf/v3/internal/bitwise";_ga "github.com/unidoc/unipdf/v3/internal/imageutil";_g "io";);type SampleWriter interface{WriteSample (_gg uint32 )error ;WriteSamples (_af []uint32 )error ;};func NewWriter (img _ga .ImageBase )*Writer {return &Writer {_ad :_a .NewWriterMSB (img .Data ),_gd :img ,_gdb :img .ColorComponents ,_gae :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_gda *Writer )WriteSamples (samples []uint32 )error {for _bbd :=0;_bbd < len (samples );_bbd ++{if _fbg :=_gda .WriteSample (samples [_bbd ]);_fbg !=nil {return _fbg ;};};return nil ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _deb []uint32 ;_fd :=bitsPerSample ;var _fe uint32 ;var _fea byte ;_gcg :=0;_ead :=0;_dd :=0;for _dd < len (data ){if _gcg > 0{_faf :=_gcg ;if _fd < _faf {_faf =_fd ;};_fe =(_fe <<uint (_faf ))|uint32 (_fea >>uint (8-_faf ));_gcg -=_faf ;if _gcg > 0{_fea =_fea <<uint (_faf );}else {_fea =0;};_fd -=_faf ;if _fd ==0{_deb =append (_deb ,_fe );_fd =bitsPerSample ;_fe =0;_ead ++;};}else {_bf :=data [_dd ];_dd ++;_baa :=8;if _fd < _baa {_baa =_fd ;};_gcg =8-_baa ;_fe =(_fe <<uint (_baa ))|uint32 (_bf >>uint (_gcg ));if _baa < 8{_fea =_bf <<uint (_baa );};_fd -=_baa ;if _fd ==0{_deb =append (_deb ,_fe );_fd =bitsPerSample ;_fe =0;_ead ++;};};};for _gcg >=bitsPerSample {_ec :=_gcg ;if _fd < _ec {_ec =_fd ;};_fe =(_fe <<uint (_ec ))|uint32 (_fea >>uint (8-_ec ));_gcg -=_ec ;if _gcg > 0{_fea =_fea <<uint (_ec );}else {_fea =0;};_fd -=_ec ;if _fd ==0{_deb =append (_deb ,_fe );_fd =bitsPerSample ;_fe =0;_ead ++;};};return _deb ;};type Writer struct{_gd _ga .ImageBase ;_ad *_a .Writer ;_cg ,_gdb int ;_gae bool ;};func (_beg *Writer )WriteSample (sample uint32 )error {if _ ,_gf :=_beg ._ad .WriteBits (uint64 (sample ),_beg ._gd .BitsPerComponent );_gf !=nil {return _gf ;};_beg ._gdb --;if _beg ._gdb ==0{_beg ._gdb =_beg ._gd .ColorComponents ;_beg ._cg ++;};if _beg ._cg ==_beg ._gd .Width {if _beg ._gae {_beg ._ad .FinishByte ();};_beg ._cg =0;};return nil ;};func (_b *Reader )ReadSamples (samples []uint32 )(_de error ){for _ba :=0;_ba < len (samples );_ba ++{samples [_ba ],_de =_b .ReadSample ();if _de !=nil {return _de ;};};return nil ;};func NewReader (img _ga .ImageBase )*Reader {return &Reader {_d :_a .NewReader (img .Data ),_fa :img ,_fac :img .ColorComponents ,_df :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};func (_db *Reader )ReadSample ()(uint32 ,error ){if _db ._ea ==_db ._fa .Height {return 0,_g .EOF ;};_ac ,_gc :=_db ._d .ReadBits (byte (_db ._fa .BitsPerComponent ));if _gc !=nil {return 0,_gc ;};_db ._fac --;if _db ._fac ==0{_db ._fac =_db ._fa .ColorComponents ;_db ._e ++;};if _db ._e ==_db ._fa .Width {if _db ._df {_db ._d .ConsumeRemainingBits ();};_db ._e =0;_db ._ea ++;};return uint32 (_ac ),nil ;};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_gag []uint32 )error ;};type Reader struct{_fa _ga .ImageBase ;_d *_a .Reader ;_e ,_ea ,_fac int ;_df bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _def []uint32 ;_fb :=bitsPerOutputSample ;var _be uint32 ;var _ge uint32 ;_dfc :=0;_c :=0;_ecc :=0;for _ecc < len (data ){if _dfc > 0{_cd :=_dfc ;if _fb < _cd {_cd =_fb ;};_be =(_be <<uint (_cd ))|(_ge >>uint (bitsPerInputSample -_cd ));_dfc -=_cd ;if _dfc > 0{_ge =_ge <<uint (_cd );}else {_ge =0;};_fb -=_cd ;if _fb ==0{_def =append (_def ,_be );_fb =bitsPerOutputSample ;_be =0;_c ++;};}else {_bb :=data [_ecc ];_ecc ++;_da :=bitsPerInputSample ;if _fb < _da {_da =_fb ;};_dfc =bitsPerInputSample -_da ;_be =(_be <<uint (_da ))|(_bb >>uint (_dfc ));if _da < bitsPerInputSample {_ge =_bb <<uint (_da );};_fb -=_da ;if _fb ==0{_def =append (_def ,_be );_fb =bitsPerOutputSample ;_be =0;_c ++;};};};for _dfc >=bitsPerOutputSample {_ed :=_dfc ;if _fb < _ed {_ed =_fb ;};_be =(_be <<uint (_ed ))|(_ge >>uint (bitsPerInputSample -_ed ));_dfc -=_ed ;if _dfc > 0{_ge =_ge <<uint (_ed );}else {_ge =0;};_fb -=_ed ;if _fb ==0{_def =append (_def ,_be );_fb =bitsPerOutputSample ;_be =0;_c ++;};};if _fb > 0&&_fb < bitsPerOutputSample {_be <<=uint (_fb );_def =append (_def ,_be );};return _def ;};