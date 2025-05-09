^XA
^PON
^LH0,0
^FWN
^FX Distributor Code
^CF0,25
^FO380,25^FD<<DistributorCode>>^FS
^FX logic for font size (replace as needed dynamically)
^CF0,<<FontSize>>
^FB900,<<TitleLines>>,,
^FO50,180^FD<<ShortDescription>>^FS
^FX Pack Size and Product Size
^CF0,70
^FO380,75^FD<<CaseQty>>x<<ProductSizeOption>><<UnitOfMeasure>>^FS
^FX UBN Code
^CF0,20
^FO220,25^FDUBN^FS
^FO220,45^BXN,4,200,26,26
^FD<<UBNCode>>^FS
^FX EPS Code
^CF0,20
^FO700,25^FDEPS^FS
^FO700,45^BXN,4,200,26,26
^FD<<EPSCode>>^FS
^FX Third section with bar code.
^FO175,325^BY5,3
^B2N,175,N,N,N
^FD<<Outer1DBarcode>>^FS
^CF0,40
^FB1000,2,,
^FO475,510^FD<<Outer1DBarcode>>^FS
^CF0,40
^FO20,550^FDLOT: <<LotNo>>^FS
^CF0,40
^FO700,550^FD<<ExpiryDate>>^FS
^XZ