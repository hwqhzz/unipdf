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

package basic ;import _g "github.com/unidoc/unipdf/v3/internal/jbig2/errors";func (_d IntsMap )Add (key uint64 ,value int ){_d [key ]=append (_d [key ],value )};func (_gf *NumSlice )Add (v float32 ){*_gf =append (*_gf ,v )};func (_ca IntsMap )GetSlice (key uint64 )([]int ,bool ){_gb ,_e :=_ca [key ];if !_e {return nil ,false ;};return _gb ,true ;};func NewIntSlice (i int )*IntSlice {_dc :=IntSlice (make ([]int ,i ));return &_dc };func (_fd IntsMap )Delete (key uint64 ){delete (_fd ,key )};func (_a *IntSlice )Add (v int )error {if _a ==nil {return _g .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");};*_a =append (*_a ,v );return nil ;};type IntsMap map[uint64 ][]int ;func NewNumSlice (i int )*NumSlice {_ef :=NumSlice (make ([]float32 ,i ));return &_ef };func (_cb *Stack )Len ()int {return len (_cb .Data )};func (_fde NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_fde )-1{return 0,_g .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};return _fde [i ],nil ;};func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};type NumSlice []float32 ;func (_ad *Stack )Peek ()(_gde interface{},_ee bool ){return _ad .peek ()};func (_ce IntSlice )Size ()int {return len (_ce )};func (_efb *Stack )top ()int {return len (_efb .Data )-1};type IntSlice []int ;func (_aa *IntSlice )Copy ()*IntSlice {_ab :=IntSlice (make ([]int ,len (*_aa )));copy (_ab ,*_aa );return &_ab ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func Max (x ,y int )int {if x > y {return x ;};return y ;};func (_dg IntsMap )Get (key uint64 )(int ,bool ){_de ,_f :=_dg [key ];if !_f {return 0,false ;};if len (_de )==0{return 0,false ;};return _de [0],true ;};func (_dd IntSlice )Get (index int )(int ,error ){if index > len (_dd )-1{return 0,_g .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );};return _dd [index ],nil ;};func (_gcb *Stack )Push (v interface{}){_gcb .Data =append (_gcb .Data ,v )};func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_efa *NumSlice )AddInt (v int ){*_efa =append (*_efa ,float32 (v ))};type Stack struct{Data []interface{};Aux *Stack ;};func Min (x ,y int )int {if x < y {return x ;};return y ;};func (_ae NumSlice )GetIntSlice ()[]int {_cc :=make ([]int ,len (_ae ));for _ge ,_gc :=range _ae {_cc [_ge ]=int (_gc );};return _cc ;};func (_gbf *Stack )Pop ()(_gcg interface{},_dge bool ){_gcg ,_dge =_gbf .peek ();if !_dge {return nil ,_dge ;};_gbf .Data =_gbf .Data [:_gbf .top ()];return _gcg ,true ;};func (_dcd NumSlice )GetInt (i int )(int ,error ){const _gd ="\u0047\u0065\u0074\u0049\u006e\u0074";if i < 0||i > len (_dcd )-1{return 0,_g .Errorf (_gd ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );};_fg :=_dcd [i ];return int (_fg +Sign (_fg )*0.5),nil ;};func (_fe *Stack )peek ()(interface{},bool ){_dda :=_fe .top ();if _dda ==-1{return nil ,false ;};return _fe .Data [_dda ],true ;};