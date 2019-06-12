unit api;

interface
uses
 uEDCode,
 System.Classes,
 System.SysUtils,
 EDCode,
 Crc,
 Vcl.Dialogs
 ;

implementation
procedure cDecodeStream_uEDCode(_instm:TStream;_outstm:TStream;_key:TStream);stdcall;
var
bykey:TBytes;
begin
 SetLength(bykey,_key.Size);
 _key.Read(bykey,_key.Size);
 uEDCode.DecodeStream(_instm,_outstm,stringof(bykey));
 _outstm.Position:=0;
end;
procedure cEncodeStream_uEDCode(_instm:TStream;_outstm:TStream;_key:TStream);stdcall;
var
bykey:TBytes;
begin
 SetLength(bykey,_key.Size);
 _key.Read(bykey,_key.Size);
 uEDCode.EncodeStream(_instm,_outstm,stringof(bykey));
 _outstm.Position:=0;
end;
procedure cSetPassWord_EDcode(_instm:TStream);stdcall;
var
pin:Ansistring;
begin
    SetLength(pin,_instm.Size);
    _instm.ReadData(pansichar(pin),_instm.Size);
    EDCode.SetPassWord(pin);
end;
function cCrc32(_instm:TStream): Cardinal;stdcall;
var
byin:TBytes;
begin
   SetLength(byin,_instm.Size);
   _instm.Read(byin[0],_instm.Size);
   Result:=Crc.Crc32(@byin[0],length(byin));
end;
procedure cEncodeString_EDcode(_instm:TStream;_outstm:TStream);stdcall;
var
r:AnsiString;
byin:TBytes;
begin
    SetLength(byin,_instm.Size);
    _instm.Read(byin,_instm.Size);
    r:=EDCode.EncodeString(StringOf(byin));
    _outstm.Write(BytesOf(r)[0],length(r));
    _outstm.Position:=0;
end;
procedure cEncodeString_uEDCode(_instm:TStream;_outstm:TStream;_key:TStream);stdcall;
var
r:AnsiString;
byin:TBytes;
bykey:TBytes;
begin
    SetLength(byin,_instm.Size);
    _instm.Read(byin[0],_instm.Size);
     SetLength(bykey,_key.Size);
    _key.Read(bykey,_key.Size);
    r:=uEDCode.EncodeString(StringOf(byin),stringof(bykey));
    _outstm.Write(BytesOf(r)[0],length(r));
    _outstm.Position:=0;
end;
procedure cDecodeString_EDcode(_instm:TStream;_outstm:TStream);stdcall;
var
r:AnsiString;
byin:TBytes;
begin
        SetLength(byin,_instm.Size);
    _instm.Read(byin[0],_instm.Size);
    r:=EDCode.DeCodeString(StringOf(byin));
    _outstm.Write(BytesOf(r)[0],length(r));
    _outstm.Position:=0;
end;
procedure cDecodeString_uEDCode(_instm:TStream;_outstm:TStream;_key:TStream);stdcall;
var
r:AnsiString;
byin:TBytes;
bykey:TBytes;
begin
    SetLength(byin,_instm.Size);
    _instm.Read(byin,_instm.Size);
    SetLength(bykey,_key.Size);
    _key.Read(bykey,_key.Size);
    r:=uEDCode.DecodeString(StringOf(byin),stringof(bykey));
    _outstm.Write(BytesOf(r)[0],length(r));
    _outstm.Position:=0;
end;
procedure cBase64DecodeEx_EDcode(_instm:TStream;_outstm:TStream;_len:Integer);stdcall;
var
r:AnsiString;
byin:TBytes;
begin
    SetLength(byin,_instm.Size);
    _instm.Read(byin,_instm.Size);
    SetLength(r,_len);
    EDCode.Base64DecodeEx(StringOf(byin),PAnsiChar(r),_len);
    _outstm.Write(BytesOf(r)[0],length(r));
    _outstm.Position:=0;
end;
procedure cDecryptAES_EDcode(_instm:TStream;_outstm:TStream);stdcall;
var
inbuf,outbuf:TAESBuffer;
begin
    _instm.Read(inbuf[0],16);
    EDCode.DecryptAES(inbuf,EDCode.FDCKey128,outbuf);
    _outstm.Write(outbuf[0],16);
    _outstm.Position:=0;
end;
procedure cEncryptAES_EDcode(_instm:TStream;_outstm:TStream);stdcall;
var
inbuf,outbuf:TAESBuffer;
begin
    _instm.Read(inbuf[0],16);
    EDCode.EncryptAES(inbuf,EDCode.FECKey128,outbuf);
    _outstm.Write(outbuf[0],16);
    _outstm.Position:=0;
end;
function BytestoHexString(ABytes: TBytes; len: Integer): AnsiString;
begin
  SetLength(Result, len*2);
  BinToHex(@ABytes[0], PAnsiChar(Result), len);
end;
procedure cBase64Encode_EDcode(_instm:TStream;_outstm:TStream;_len:Integer);stdcall;
var
r:AnsiString;
byin:TBytes;
begin
    SetLength(byin,_instm.Size);
    _instm.Read(byin,_instm.Size);
    SetLength(r,_len);
    EDCode.Base64Encode(pansichar(pbyte(@byin[0])),_len,r);
    _outstm.Write(BytesOf(r)[0],_len);
    _outstm.Position:=0;
end;
exports
cDecodeStream_uEDCode,
cEncodeStream_uEDCode,
cSetPassWord_EDcode,
cEncodeString_EDcode,
cEncodeString_uEDCode,
cDecodeString_EDcode,
cDecodeString_uEDCode,
cBase64DecodeEx_EDcode,
cDecryptAES_EDcode,
cEncryptAES_EDcode,
cBase64Encode_EDcode,
cCrc32
;
end.
