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

// Package draw has handy features for defining paths which can be used to draw content on a PDF page.  Handles
// defining paths as points, vector calculations and conversion to PDF content stream data which can be used in
// page content streams and XObject forms and thus also in annotation appearance streams.
//
// Also defines utility functions for drawing common shapes such as rectangles, lines and circles (ovals).
package draw ;import (_d "fmt";_ff "github.com/unidoc/unipdf/v3/contentstream";_db "github.com/unidoc/unipdf/v3/core";_a "github.com/unidoc/unipdf/v3/internal/transform";_b "github.com/unidoc/unipdf/v3/model";_f "math";);

// Flip changes the sign of the vector: -vector.
func (_dba Vector )Flip ()Vector {_dgac :=_dba .Magnitude ();_bfc :=_dba .GetPolarAngle ();_dba .Dx =_dgac *_f .Cos (_bfc +_f .Pi );_dba .Dy =_dgac *_f .Sin (_bfc +_f .Pi );return _dba ;};

// NewCubicBezierPath returns a new empty cubic Bezier path.
func NewCubicBezierPath ()CubicBezierPath {_fe :=CubicBezierPath {};_fe .Curves =[]CubicBezierCurve {};return _fe ;};

// Add adds the specified vector to the current one and returns the result.
func (_fae Vector )Add (other Vector )Vector {_fae .Dx +=other .Dx ;_fae .Dy +=other .Dy ;return _fae };

// Scale scales the vector by the specified factor.
func (_geaf Vector )Scale (factor float64 )Vector {_adf :=_geaf .Magnitude ();_feec :=_geaf .GetPolarAngle ();_geaf .Dx =factor *_adf *_f .Cos (_feec );_geaf .Dy =factor *_adf *_f .Sin (_feec );return _geaf ;};const (LineStyleSolid LineStyle =0;LineStyleDashed LineStyle =1;);

// GetBoundingBox returns the bounding box of the path.
func (_gd Path )GetBoundingBox ()BoundingBox {_fgb :=BoundingBox {};_gdb :=0.0;_dd :=0.0;_dgf :=0.0;_cd :=0.0;for _dcce ,_efg :=range _gd .Points {if _dcce ==0{_gdb =_efg .X ;_dd =_efg .X ;_dgf =_efg .Y ;_cd =_efg .Y ;continue ;};if _efg .X < _gdb {_gdb =_efg .X ;};if _efg .X > _dd {_dd =_efg .X ;};if _efg .Y < _dgf {_dgf =_efg .Y ;};if _efg .Y > _cd {_cd =_efg .Y ;};};_fgb .X =_gdb ;_fgb .Y =_dgf ;_fgb .Width =_dd -_gdb ;_fgb .Height =_cd -_dgf ;return _fgb ;};

// Polyline defines a slice of points that are connected as straight lines.
type Polyline struct{Points []Point ;LineColor *_b .PdfColorDeviceRGB ;LineWidth float64 ;};

// NewPoint returns a new point with the coordinates x, y.
func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};

// GetBoundingBox returns the bounding box of the Bezier path.
func (_af CubicBezierPath )GetBoundingBox ()Rectangle {_afc :=Rectangle {};_fg :=0.0;_fac :=0.0;_aa :=0.0;_fd :=0.0;for _gf ,_ed :=range _af .Curves {_dc :=_ed .GetBounds ();if _gf ==0{_fg =_dc .Llx ;_fac =_dc .Urx ;_aa =_dc .Lly ;_fd =_dc .Ury ;continue ;};if _dc .Llx < _fg {_fg =_dc .Llx ;};if _dc .Urx > _fac {_fac =_dc .Urx ;};if _dc .Lly < _aa {_aa =_dc .Lly ;};if _dc .Ury > _fd {_fd =_dc .Ury ;};};_afc .X =_fg ;_afc .Y =_aa ;_afc .Width =_fac -_fg ;_afc .Height =_fd -_aa ;return _afc ;};const (LineEndingStyleNone LineEndingStyle =0;LineEndingStyleArrow LineEndingStyle =1;LineEndingStyleButt LineEndingStyle =2;);

// Magnitude returns the magnitude of the vector.
func (_eeb Vector )Magnitude ()float64 {return _f .Sqrt (_f .Pow (_eeb .Dx ,2.0)+_f .Pow (_eeb .Dy ,2.0))};

// GetBounds returns the bounding box of the Bezier curve.
func (_ee CubicBezierCurve )GetBounds ()_b .PdfRectangle {_bc :=_ee .P0 .X ;_ef :=_ee .P0 .X ;_be :=_ee .P0 .Y ;_ga :=_ee .P0 .Y ;for _de :=0.0;_de <=1.0;_de +=0.001{Rx :=_ee .P0 .X *_f .Pow (1-_de ,3)+_ee .P1 .X *3*_de *_f .Pow (1-_de ,2)+_ee .P2 .X *3*_f .Pow (_de ,2)*(1-_de )+_ee .P3 .X *_f .Pow (_de ,3);Ry :=_ee .P0 .Y *_f .Pow (1-_de ,3)+_ee .P1 .Y *3*_de *_f .Pow (1-_de ,2)+_ee .P2 .Y *3*_f .Pow (_de ,2)*(1-_de )+_ee .P3 .Y *_f .Pow (_de ,3);if Rx < _bc {_bc =Rx ;};if Rx > _ef {_ef =Rx ;};if Ry < _be {_be =Ry ;};if Ry > _ga {_ga =Ry ;};};_df :=_b .PdfRectangle {};_df .Llx =_bc ;_df .Lly =_be ;_df .Urx =_ef ;_df .Ury =_ga ;return _df ;};

// PolyBezierCurve represents a composite curve that is the result of
// joining multiple cubic Bezier curves.
type PolyBezierCurve struct{Curves []CubicBezierCurve ;BorderWidth float64 ;BorderColor *_b .PdfColorDeviceRGB ;FillEnabled bool ;FillColor *_b .PdfColorDeviceRGB ;};

// NewVectorBetween returns a new vector with the direction specified by
// the subtraction of point a from point b (b-a).
func NewVectorBetween (a Point ,b Point )Vector {_fef :=Vector {};_fef .Dx =b .X -a .X ;_fef .Dy =b .Y -a .Y ;return _fef ;};

// BasicLine defines a line between point 1 (X1,Y1) and point 2 (X2,Y2). The line has a specified width, color and opacity.
type BasicLine struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_b .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineStyle LineStyle ;};

// AddVector adds vector to a point.
func (_ag Point )AddVector (v Vector )Point {_ag .X +=v .Dx ;_ag .Y +=v .Dy ;return _ag };

// FlipX flips the sign of the Dx component of the vector.
func (_bea Vector )FlipX ()Vector {_bea .Dx =-_bea .Dx ;return _bea };

// Circle represents a circle shape with fill and border properties that can be drawn to a PDF content stream.
type Circle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_b .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_b .PdfColorDeviceRGB ;Opacity float64 ;};

// FlipY flips the sign of the Dy component of the vector.
func (_feb Vector )FlipY ()Vector {_feb .Dy =-_feb .Dy ;return _feb };

// Draw draws the circle. Can specify a graphics state (gsName) for setting opacity etc.  Otherwise leave empty ("").
// Returns the content stream as a byte array, the bounding box and an error on failure.
func (_cf Circle )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_eecd :=_cf .Width /2;_fcf :=_cf .Height /2;if _cf .BorderEnabled {_eecd -=_cf .BorderWidth /2;_fcf -=_cf .BorderWidth /2;};_dccb :=0.551784;_edg :=_eecd *_dccb ;_ddb :=_fcf *_dccb ;_bbd :=NewCubicBezierPath ();_bbd =_bbd .AppendCurve (NewCubicBezierCurve (-_eecd ,0,-_eecd ,_ddb ,-_edg ,_fcf ,0,_fcf ));_bbd =_bbd .AppendCurve (NewCubicBezierCurve (0,_fcf ,_edg ,_fcf ,_eecd ,_ddb ,_eecd ,0));_bbd =_bbd .AppendCurve (NewCubicBezierCurve (_eecd ,0,_eecd ,-_ddb ,_edg ,-_fcf ,0,-_fcf ));_bbd =_bbd .AppendCurve (NewCubicBezierCurve (0,-_fcf ,-_edg ,-_fcf ,-_eecd ,-_ddb ,-_eecd ,0));_bbd =_bbd .Offset (_eecd ,_fcf );if _cf .BorderEnabled {_bbd =_bbd .Offset (_cf .BorderWidth /2,_cf .BorderWidth /2);};if _cf .X !=0||_cf .Y !=0{_bbd =_bbd .Offset (_cf .X ,_cf .Y );};_dgc :=_ff .NewContentCreator ();_dgc .Add_q ();if _cf .FillEnabled {_dgc .Add_rg (_cf .FillColor .R (),_cf .FillColor .G (),_cf .FillColor .B ());};if _cf .BorderEnabled {_dgc .Add_RG (_cf .BorderColor .R (),_cf .BorderColor .G (),_cf .BorderColor .B ());_dgc .Add_w (_cf .BorderWidth );};if len (gsName )> 1{_dgc .Add_gs (_db .PdfObjectName (gsName ));};DrawBezierPathWithCreator (_bbd ,_dgc );_dgc .Add_h ();if _cf .FillEnabled &&_cf .BorderEnabled {_dgc .Add_B ();}else if _cf .FillEnabled {_dgc .Add_f ();}else if _cf .BorderEnabled {_dgc .Add_S ();};_dgc .Add_Q ();_eed :=_bbd .GetBoundingBox ();if _cf .BorderEnabled {_eed .Height +=_cf .BorderWidth ;_eed .Width +=_cf .BorderWidth ;_eed .X -=_cf .BorderWidth /2;_eed .Y -=_cf .BorderWidth /2;};return _dgc .Bytes (),_eed .ToPdfRectangle (),nil ;};

// Offset shifts the path with the specified offsets.
func (_fc Path )Offset (offX ,offY float64 )Path {for _ece ,_bb :=range _fc .Points {_fc .Points [_ece ]=_bb .Add (offX ,offY );};return _fc ;};

// Rotate rotates the vector by the specified angle.
func (_afbbf Vector )Rotate (phi float64 )Vector {_bcd :=_afbbf .Magnitude ();_cebgc :=_afbbf .GetPolarAngle ();return NewVectorPolar (_bcd ,_cebgc +phi );};

// LineEndingStyle defines the line ending style for lines.
// The currently supported line ending styles are None, Arrow (ClosedArrow) and Butt.
type LineEndingStyle int ;

// NewPath returns a new empty path.
func NewPath ()Path {return Path {}};

// Offset shifts the Bezier path with the specified offsets.
func (_fa CubicBezierPath )Offset (offX ,offY float64 )CubicBezierPath {for _gb ,_gg :=range _fa .Curves {_fa .Curves [_gb ]=_gg .AddOffsetXY (offX ,offY );};return _fa ;};

// Point represents a two-dimensional point.
type Point struct{X float64 ;Y float64 ;};

// ToPdfRectangle returns the rectangle as a PDF rectangle.
func (_fcfg Rectangle )ToPdfRectangle ()*_b .PdfRectangle {return &_b .PdfRectangle {Llx :_fcfg .X ,Lly :_fcfg .Y ,Urx :_fcfg .X +_fcfg .Width ,Ury :_fcfg .Y +_fcfg .Height };};

// LineStyle refers to how the line will be created.
type LineStyle int ;

// AddOffsetXY adds X,Y offset to all points on a curve.
func (_e CubicBezierCurve )AddOffsetXY (offX ,offY float64 )CubicBezierCurve {_e .P0 .X +=offX ;_e .P1 .X +=offX ;_e .P2 .X +=offX ;_e .P3 .X +=offX ;_e .P0 .Y +=offY ;_e .P1 .Y +=offY ;_e .P2 .Y +=offY ;_e .P3 .Y +=offY ;return _e ;};

// Draw draws the polygon. A graphics state name can be specified for
// setting the polygon properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polygon
// bounding box.
func (_fcc Polygon )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_aca :=_ff .NewContentCreator ();_aca .Add_q ();_fcc .FillEnabled =_fcc .FillEnabled &&_fcc .FillColor !=nil ;if _fcc .FillEnabled {_aca .Add_rg (_fcc .FillColor .R (),_fcc .FillColor .G (),_fcc .FillColor .B ());};_fcc .BorderEnabled =_fcc .BorderEnabled &&_fcc .BorderColor !=nil ;if _fcc .BorderEnabled {_aca .Add_RG (_fcc .BorderColor .R (),_fcc .BorderColor .G (),_fcc .BorderColor .B ());_aca .Add_w (_fcc .BorderWidth );};if len (gsName )> 1{_aca .Add_gs (_db .PdfObjectName (gsName ));};_gag :=NewPath ();for _ ,_cbe :=range _fcc .Points {for _aga ,_ceb :=range _cbe {_gag =_gag .AppendPoint (_ceb );if _aga ==0{_aca .Add_m (_ceb .X ,_ceb .Y );}else {_aca .Add_l (_ceb .X ,_ceb .Y );};};_aca .Add_h ();};if _fcc .FillEnabled &&_fcc .BorderEnabled {_aca .Add_B ();}else if _fcc .FillEnabled {_aca .Add_f ();}else if _fcc .BorderEnabled {_aca .Add_S ();};_aca .Add_Q ();return _aca .Bytes (),_gag .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Path consists of straight line connections between each point defined in an array of points.
type Path struct{Points []Point ;};

// AppendPoint adds the specified point to the path.
func (_dcc Path )AppendPoint (point Point )Path {_dcc .Points =append (_dcc .Points ,point );return _dcc };

// Draw draws the line to PDF contentstream. Generates the content stream which can be used in page contents or
// appearance stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error
// if one occurred.
func (_ad Line )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_fgf ,_gdbb :=_ad .X1 ,_ad .X2 ;_fdf ,_agd :=_ad .Y1 ,_ad .Y2 ;_gdbg :=_agd -_fdf ;_gda :=_gdbb -_fgf ;_ede :=_f .Atan2 (_gdbg ,_gda );L :=_f .Sqrt (_f .Pow (_gda ,2.0)+_f .Pow (_gdbg ,2.0));_fdc :=_ad .LineWidth ;_agc :=_f .Pi ;_fgfa :=1.0;if _gda < 0{_fgfa *=-1.0;};if _gdbg < 0{_fgfa *=-1.0;};VsX :=_fgfa *(-_fdc /2*_f .Cos (_ede +_agc /2));VsY :=_fgfa *(-_fdc /2*_f .Sin (_ede +_agc /2)+_fdc *_f .Sin (_ede +_agc /2));V1X :=VsX +_fdc /2*_f .Cos (_ede +_agc /2);V1Y :=VsY +_fdc /2*_f .Sin (_ede +_agc /2);V2X :=VsX +_fdc /2*_f .Cos (_ede +_agc /2)+L *_f .Cos (_ede );V2Y :=VsY +_fdc /2*_f .Sin (_ede +_agc /2)+L *_f .Sin (_ede );V3X :=VsX +_fdc /2*_f .Cos (_ede +_agc /2)+L *_f .Cos (_ede )+_fdc *_f .Cos (_ede -_agc /2);V3Y :=VsY +_fdc /2*_f .Sin (_ede +_agc /2)+L *_f .Sin (_ede )+_fdc *_f .Sin (_ede -_agc /2);V4X :=VsX +_fdc /2*_f .Cos (_ede -_agc /2);V4Y :=VsY +_fdc /2*_f .Sin (_ede -_agc /2);_ca :=NewPath ();_ca =_ca .AppendPoint (NewPoint (V1X ,V1Y ));_ca =_ca .AppendPoint (NewPoint (V2X ,V2Y ));_ca =_ca .AppendPoint (NewPoint (V3X ,V3Y ));_ca =_ca .AppendPoint (NewPoint (V4X ,V4Y ));_dfaa :=_ad .LineEndingStyle1 ;_bd :=_ad .LineEndingStyle2 ;_bae :=3*_fdc ;_cba :=3*_fdc ;_ebd :=(_cba -_fdc )/2;if _bd ==LineEndingStyleArrow {_feg :=_ca .GetPointNumber (2);_def :=NewVectorPolar (_bae ,_ede +_agc );_ceba :=_feg .AddVector (_def );_gef :=NewVectorPolar (_cba /2,_ede +_agc /2);_fdb :=NewVectorPolar (_bae ,_ede );_afb :=NewVectorPolar (_ebd ,_ede +_agc /2);_afbb :=_ceba .AddVector (_afb );_ffg :=_fdb .Add (_gef .Flip ());_eff :=_afbb .AddVector (_ffg );_fge :=_gef .Scale (2).Flip ().Add (_ffg .Flip ());_aee :=_eff .AddVector (_fge );_bg :=_ceba .AddVector (NewVectorPolar (_fdc ,_ede -_agc /2));_cfa :=NewPath ();_cfa =_cfa .AppendPoint (_ca .GetPointNumber (1));_cfa =_cfa .AppendPoint (_ceba );_cfa =_cfa .AppendPoint (_afbb );_cfa =_cfa .AppendPoint (_eff );_cfa =_cfa .AppendPoint (_aee );_cfa =_cfa .AppendPoint (_bg );_cfa =_cfa .AppendPoint (_ca .GetPointNumber (4));_ca =_cfa ;};if _dfaa ==LineEndingStyleArrow {_eda :=_ca .GetPointNumber (1);_afe :=_ca .GetPointNumber (_ca .Length ());_adb :=NewVectorPolar (_fdc /2,_ede +_agc +_agc /2);_dcd :=_eda .AddVector (_adb );_eag :=NewVectorPolar (_bae ,_ede ).Add (NewVectorPolar (_cba /2,_ede +_agc /2));_facb :=_dcd .AddVector (_eag );_dbf :=NewVectorPolar (_ebd ,_ede -_agc /2);_fea :=_facb .AddVector (_dbf );_fee :=NewVectorPolar (_bae ,_ede );_afa :=_afe .AddVector (_fee );_aac :=NewVectorPolar (_ebd ,_ede +_agc +_agc /2);_cbf :=_afa .AddVector (_aac );_bgf :=_dcd ;_cebg :=NewPath ();_cebg =_cebg .AppendPoint (_dcd );_cebg =_cebg .AppendPoint (_facb );_cebg =_cebg .AppendPoint (_fea );for _ ,_gaab :=range _ca .Points [1:len (_ca .Points )-1]{_cebg =_cebg .AppendPoint (_gaab );};_cebg =_cebg .AppendPoint (_afa );_cebg =_cebg .AppendPoint (_cbf );_cebg =_cebg .AppendPoint (_bgf );_ca =_cebg ;};_eg :=_ff .NewContentCreator ();_eg .Add_q ().Add_rg (_ad .LineColor .R (),_ad .LineColor .G (),_ad .LineColor .B ());if len (gsName )> 1{_eg .Add_gs (_db .PdfObjectName (gsName ));};_ca =_ca .Offset (_ad .X1 ,_ad .Y1 );_cae :=_ca .GetBoundingBox ();DrawPathWithCreator (_ca ,_eg );if _ad .LineStyle ==LineStyleDashed {_eg .Add_d ([]int64 {1,1},0).Add_S ().Add_f ().Add_Q ();}else {_eg .Add_f ().Add_Q ();};return _eg .Bytes (),_cae .ToPdfRectangle (),nil ;};

// DrawPathWithCreator makes the path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawPathWithCreator (path Path ,creator *_ff .ContentCreator ){for _cgf ,_dged :=range path .Points {if _cgf ==0{creator .Add_m (_dged .X ,_dged .Y );}else {creator .Add_l (_dged .X ,_dged .Y );};};};

// NewVectorPolar returns a new vector calculated from the specified
// magnitude and angle.
func NewVectorPolar (length float64 ,theta float64 )Vector {_ebff :=Vector {};_ebff .Dx =length *_f .Cos (theta );_ebff .Dy =length *_f .Sin (theta );return _ebff ;};

// Copy returns a clone of the Bezier path.
func (_dg CubicBezierPath )Copy ()CubicBezierPath {_ec :=CubicBezierPath {};_ec .Curves =append (_ec .Curves ,_dg .Curves ...);return _ec ;};

// Rectangle is a shape with a specified Width and Height and a lower left corner at (X,Y) that can be
// drawn to a PDF content stream.  The rectangle can optionally have a border and a filling color.
// The Width/Height includes the border (if any specified), i.e. is positioned inside.
type Rectangle struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;FillEnabled bool ;FillColor *_b .PdfColorDeviceRGB ;BorderEnabled bool ;BorderWidth float64 ;BorderColor *_b .PdfColorDeviceRGB ;Opacity float64 ;};

// BoundingBox represents the smallest rectangular area that encapsulates an object.
type BoundingBox struct{X float64 ;Y float64 ;Width float64 ;Height float64 ;};func (_ge Point )String ()string {return _d .Sprintf ("(\u0025\u002e\u0031\u0066\u002c\u0025\u002e\u0031\u0066\u0029",_ge .X ,_ge .Y );};

// GetPolarAngle returns the angle the magnitude of the vector forms with the
// positive X-axis going counterclockwise.
func (_aea Vector )GetPolarAngle ()float64 {return _f .Atan2 (_aea .Dy ,_aea .Dx )};

// Draw draws the rectangle. Can specify a graphics state (gsName) for setting opacity etc.
// Otherwise leave empty (""). Returns the content stream as a byte array, bounding box and an error on failure.
func (_ced Rectangle )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_bf :=NewPath ();_bf =_bf .AppendPoint (NewPoint (0,0));_bf =_bf .AppendPoint (NewPoint (0,_ced .Height ));_bf =_bf .AppendPoint (NewPoint (_ced .Width ,_ced .Height ));_bf =_bf .AppendPoint (NewPoint (_ced .Width ,0));_bf =_bf .AppendPoint (NewPoint (0,0));if _ced .X !=0||_ced .Y !=0{_bf =_bf .Offset (_ced .X ,_ced .Y );};_ffc :=_ff .NewContentCreator ();_ffc .Add_q ();if _ced .FillEnabled {_ffc .Add_rg (_ced .FillColor .R (),_ced .FillColor .G (),_ced .FillColor .B ());};if _ced .BorderEnabled {_ffc .Add_RG (_ced .BorderColor .R (),_ced .BorderColor .G (),_ced .BorderColor .B ());_ffc .Add_w (_ced .BorderWidth );};if len (gsName )> 1{_ffc .Add_gs (_db .PdfObjectName (gsName ));};DrawPathWithCreator (_bf ,_ffc );_ffc .Add_h ();if _ced .FillEnabled &&_ced .BorderEnabled {_ffc .Add_B ();}else if _ced .FillEnabled {_ffc .Add_f ();}else if _ced .BorderEnabled {_ffc .Add_S ();};_ffc .Add_Q ();return _ffc .Bytes (),_bf .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Polygon is a multi-point shape that can be drawn to a PDF content stream.
type Polygon struct{Points [][]Point ;FillEnabled bool ;FillColor *_b .PdfColorDeviceRGB ;BorderEnabled bool ;BorderColor *_b .PdfColorDeviceRGB ;BorderWidth float64 ;};

// Copy returns a clone of the path.
func (_gbb Path )Copy ()Path {_dfa :=Path {};_dfa .Points =append (_dfa .Points ,_gbb .Points ...);return _dfa ;};

// CubicBezierPath represents a collection of cubic Bezier curves.
type CubicBezierPath struct{Curves []CubicBezierCurve ;};

// Draw draws the polyline. A graphics state name can be specified for
// setting the polyline properties (e.g. setting the opacity). Otherwise leave
// empty (""). Returns the content stream as a byte array and the polyline
// bounding box.
func (_aef Polyline )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){if _aef .LineColor ==nil {_aef .LineColor =_b .NewPdfColorDeviceRGB (0,0,0);};_age :=NewPath ();for _ ,_cdd :=range _aef .Points {_age =_age .AppendPoint (_cdd );};_fga :=_ff .NewContentCreator ();_fga .Add_q ();_fga .Add_RG (_aef .LineColor .R (),_aef .LineColor .G (),_aef .LineColor .B ());_fga .Add_w (_aef .LineWidth );if len (gsName )> 1{_fga .Add_gs (_db .PdfObjectName (gsName ));};DrawPathWithCreator (_age ,_fga );_fga .Add_S ();_fga .Add_Q ();return _fga .Bytes (),_age .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Vector represents a two-dimensional vector.
type Vector struct{Dx float64 ;Dy float64 ;};

// DrawBezierPathWithCreator makes the bezier path with the content creator.
// Adds the PDF commands to draw the path to the creator instance.
func DrawBezierPathWithCreator (bpath CubicBezierPath ,creator *_ff .ContentCreator ){for _ggb ,_acag :=range bpath .Curves {if _ggb ==0{creator .Add_m (_acag .P0 .X ,_acag .P0 .Y );};creator .Add_c (_acag .P1 .X ,_acag .P1 .Y ,_acag .P2 .X ,_acag .P2 .Y ,_acag .P3 .X ,_acag .P3 .Y );};};

// AppendCurve appends the specified Bezier curve to the path.
func (_ba CubicBezierPath )AppendCurve (curve CubicBezierCurve )CubicBezierPath {_ba .Curves =append (_ba .Curves ,curve );return _ba ;};

// Line defines a line shape between point 1 (X1,Y1) and point 2 (X2,Y2).  The line ending styles can be none (regular line),
// or arrows at either end.  The line also has a specified width, color and opacity.
type Line struct{X1 float64 ;Y1 float64 ;X2 float64 ;Y2 float64 ;LineColor *_b .PdfColorDeviceRGB ;Opacity float64 ;LineWidth float64 ;LineEndingStyle1 LineEndingStyle ;LineEndingStyle2 LineEndingStyle ;LineStyle LineStyle ;};

// Draw draws the composite Bezier curve. A graphics state name can be
// specified for setting the curve properties (e.g. setting the opacity).
// Otherwise leave empty (""). Returns the content stream as a byte array and
// the curve bounding box.
func (_aad PolyBezierCurve )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){if _aad .BorderColor ==nil {_aad .BorderColor =_b .NewPdfColorDeviceRGB (0,0,0);};_dgg :=NewCubicBezierPath ();for _ ,_ce :=range _aad .Curves {_dgg =_dgg .AppendCurve (_ce );};_fgg :=_ff .NewContentCreator ();_fgg .Add_q ();_aad .FillEnabled =_aad .FillEnabled &&_aad .FillColor !=nil ;if _aad .FillEnabled {_fgg .Add_rg (_aad .FillColor .R (),_aad .FillColor .G (),_aad .FillColor .B ());};_fgg .Add_RG (_aad .BorderColor .R (),_aad .BorderColor .G (),_aad .BorderColor .B ());_fgg .Add_w (_aad .BorderWidth );if len (gsName )> 1{_fgg .Add_gs (_db .PdfObjectName (gsName ));};for _ ,_dga :=range _dgg .Curves {_fgg .Add_m (_dga .P0 .X ,_dga .P0 .Y );_fgg .Add_c (_dga .P1 .X ,_dga .P1 .Y ,_dga .P2 .X ,_dga .P2 .Y ,_dga .P3 .X ,_dga .P3 .Y );};if _aad .FillEnabled {_fgg .Add_h ();_fgg .Add_B ();}else {_fgg .Add_S ();};_fgg .Add_Q ();return _fgg .Bytes (),_dgg .GetBoundingBox ().ToPdfRectangle (),nil ;};

// Length returns the number of points in the path.
func (_ac Path )Length ()int {return len (_ac .Points )};

// Rotate returns a new Point at `p` rotated by `theta` degrees.
func (_gaa Point )Rotate (theta float64 )Point {_ae :=_a .NewPoint (_gaa .X ,_gaa .Y ).Rotate (theta );return NewPoint (_ae .X ,_ae .Y );};

// GetPointNumber returns the path point at the index specified by number.
// The index is 1-based.
func (_gga Path )GetPointNumber (number int )Point {if number < 1||number > len (_gga .Points ){return Point {};};return _gga .Points [number -1];};

// ToPdfRectangle returns the bounding box as a PDF rectangle.
func (_eec BoundingBox )ToPdfRectangle ()*_b .PdfRectangle {return &_b .PdfRectangle {Llx :_eec .X ,Lly :_eec .Y ,Urx :_eec .X +_eec .Width ,Ury :_eec .Y +_eec .Height };};

// Add shifts the coordinates of the point with dx, dy and returns the result.
func (_cb Point )Add (dx ,dy float64 )Point {_cb .X +=dx ;_cb .Y +=dy ;return _cb };

// NewCubicBezierCurve returns a new cubic Bezier curve.
func NewCubicBezierCurve (x0 ,y0 ,x1 ,y1 ,x2 ,y2 ,x3 ,y3 float64 )CubicBezierCurve {_c :=CubicBezierCurve {};_c .P0 =NewPoint (x0 ,y0 );_c .P1 =NewPoint (x1 ,y1 );_c .P2 =NewPoint (x2 ,y2 );_c .P3 =NewPoint (x3 ,y3 );return _c ;};

// Draw draws the basic line to PDF. Generates the content stream which can be used in page contents or appearance
// stream of annotation. Returns the stream content, XForm bounding box (local), bounding box and an error if
// one occurred.
func (_ebf BasicLine )Draw (gsName string )([]byte ,*_b .PdfRectangle ,error ){_gea :=_ebf .LineWidth ;_cc :=NewPath ();_cc =_cc .AppendPoint (NewPoint (_ebf .X1 ,_ebf .Y1 ));_cc =_cc .AppendPoint (NewPoint (_ebf .X2 ,_ebf .Y2 ));_geae :=_ff .NewContentCreator ();_cff :=_cc .GetBoundingBox ();DrawPathWithCreator (_cc ,_geae );if _ebf .LineStyle ==LineStyleDashed {_geae .Add_d ([]int64 {1,1},0);};_geae .Add_RG (_ebf .LineColor .R (),_ebf .LineColor .G (),_ebf .LineColor .B ()).Add_w (_gea ).Add_S ().Add_Q ();return _geae .Bytes (),_cff .ToPdfRectangle (),nil ;};

// RemovePoint removes the point at the index specified by number from the
// path. The index is 1-based.
func (_eb Path )RemovePoint (number int )Path {if number < 1||number > len (_eb .Points ){return _eb ;};_ea :=number -1;_eb .Points =append (_eb .Points [:_ea ],_eb .Points [_ea +1:]...);return _eb ;};

// CubicBezierCurve is defined by:
// R(t) = P0*(1-t)^3 + P1*3*t*(1-t)^2 + P2*3*t^2*(1-t) + P3*t^3
// where P0 is the current point, P1, P2 control points and P3 the final point.
type CubicBezierCurve struct{P0 Point ;P1 Point ;P2 Point ;P3 Point ;};

// NewVector returns a new vector with the direction specified by dx and dy.
func NewVector (dx ,dy float64 )Vector {_adbb :=Vector {};_adbb .Dx =dx ;_adbb .Dy =dy ;return _adbb };