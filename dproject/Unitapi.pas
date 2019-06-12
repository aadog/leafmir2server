unit Unitapi;

interface
uses
SysUtils,
Classes,
uEDCode,
EDCode,
Crc;

implementation

procedure cDecodeStream_uEDCode(_instm:pbyte;_inlen:Integer;_outstm:pbyte;var _outlen:integer;_key:PAnsiChar);stdcall;
var
tmpbuf:Tbytes;
key:string;
instm:TMemoryStream;
outstm:TMemoryStream;
begin

    instm:=TMemoryStream.create;
    outstm:=TMemoryStream.Create;
    try

       instm.WriteBuffer(_instm^,_inlen);
       instm.Position:=0;
       uEDCode.DecodeStream(instm,outstm,stringof(bytesof(_key)));
       outstm.Seek(0, soFromBeginning);
       if _outstm<>nil then
       begin
           Move(outstm.Memory^, _outstm^, outstm.Size);
       end;
       _outlen := outstm.Size;
    finally
      instm.Free;
      outstm.Free;
    end;
end;
procedure cEncodeStream_uEDCode(_instm:pbyte;_inlen:Integer;_outstm:pbyte;var _outlen:integer;_key:PAnsiChar);stdcall;
var
instm:TMemoryStream;
outstm:TMemoryStream;
begin
    instm:=TMemoryStream.create;
    outstm:=TMemoryStream.Create;
    try
       instm.WriteBuffer(_instm^,_inlen);
       instm.Position:=0;
       uEDCode.EncodeStream(instm,outstm,_key);
       outstm.Seek(0, soFromBeginning);
       if _outstm<>nil then
       begin
           Move(outstm.Memory^, _outstm^, outstm.Size);
       end;
       _outlen := outstm.Size;
    finally
      instm.Free;
      outstm.Free;
    end;
end;
procedure cEncodeString_uEDCode(_in:pbyte;_key:PAnsiChar;_outbuf:pbyte;var _outlen:Integer);stdcall;
var
r:AnsiString;
begin

    r:=uEDCode.EncodeString(pansichar(_in),_key);
    if _outbuf<>nil then
    begin
       Move(bytesof(r)[0],_outbuf^,Length(r));
    end;
    _outlen:=Length(r);
end;
procedure cDecodeString_uEDCode(_in:pbyte;_key:PAnsiChar;_outbuf:pbyte;var _outlen:Integer);stdcall;
var
r:AnsiString;
begin
    r:=uEDCode.DecodeString(PAnsiChar(_in),_key);
    if _outbuf<>nil then
    begin
       Move(bytesof(r)[0],_outbuf^,Length(r));
    end;
    _outlen:=Length(r);
end;
procedure cEncodeString_EDcode(_in:pbyte;_outbuf:pbyte;var _outlen:Integer);stdcall;
var
r:AnsiString;
begin
    r:=EDCode.EncodeString(pansichar(_in));
    if _outbuf<>nil then
    begin
       Move(bytesof(r)[0],_outbuf^,Length(r));
    end;
    _outlen:=Length(r);
end;
procedure cBase64DecodeEx_EDcode(_in:pbyte;_outbuf:pbyte;_len:Integer);stdcall;
var
r:AnsiString;
begin
    SetLength(r,_len);
    EDCode.Base64DecodeEx(pansichar(_in),PAnsiChar(r),_len);
    Move(bytesof(r)[0],_outbuf^,_len);
end;
procedure cBase64Encode_EDcode(_in:pbyte;_outbuf:pbyte;_len:Integer);stdcall;
var
r:AnsiString;
begin
    SetLength(r,_len);
    EDCode.Base64Encode(pansichar(_in),_len,r);
    Move(bytesof(r)[0],_outbuf^,_len);
end;
procedure cDecodeString_EDcode(_in:pbyte;_outbuf:pbyte;var _outlen:Integer);stdcall;
var
r:AnsiString;
begin

    r:=EDCode.DeCodeString(pansichar(_in));
    if _outbuf<>nil then
    begin
       Move(bytesof(r)[0],_outbuf^,Length(r));
    end;
    _outlen:=Length(r);
end;
procedure cDecryptAES_EDcode(_in:PByte;_outbuf:PByte);stdcall;
var
inbuf,outbuf:TAESBuffer;
begin
    Move(_in^,inbuf,16);
    EDCode.DecryptAES(inbuf,EDCode.FDCKey128,outbuf);
    Move(outbuf,_outbuf^,16);
end;
procedure cEncryptAES_EDcode(_in:PByte;_outbuf:PByte);stdcall;
var
inbuf,outbuf:TAESBuffer;
begin
    Move(_in^,inbuf,16);
    EDCode.EncryptAES(inbuf,EDCode.FECKey128,outbuf);
    Move(outbuf,_outbuf^,16);
end;
procedure cSetPassWord_EDcode(_in:PByte);stdcall;
begin
    EDCode.SetPassWord(PAnsiChar(_in));
end;
function cCrc32(Buf: PByte; Len: Integer): Cardinal;stdcall;
begin
   Result:=Crc.Crc32(Buf,Len);
end;
exports
cEncodeString_uEDCode,
cDecodeString_uEDCode,
cEncodeStream_uEDCode,
cDecodeStream_uEDCode,
cEncodeString_EDcode,
cDecodeString_EDcode,
cBase64DecodeEx_EDcode,
cSetPassWord_EDcode,
cDecryptAES_EDcode,
cEncryptAES_EDcode,
cBase64Encode_EDcode,
cCrc32;


end.
