unit Unitapi;

interface
uses
Unitcrypto,SysUtils,Classes;

implementation

type PTMemoryStream=^TMemoryStream;

function StrToHex(AStr: string): string;
var
i : Integer;
ch:char;
begin

  Result:='';
  for i:=1 to length(AStr)  do
  begin
    ch:=AStr[i];
    Result:=Result+IntToHex(Ord(ch),2);
  end;
end;


function StrPasEx(ABuff: PAnsiChar; ALength: Integer): AnsiString; inline;
begin
  SetLength(Result, ALength);
  Move(ABuff^, Result[1], ALength);
end;
procedure cBase64Decode(_in:PAnsiChar;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;
var
r:PAnsiChar;
rs:AnsiString;
a1,a2:Integer;
begin
    try
      Base64Decode(_in,r,a1,a2);
      rs := StrPasEx(r,a1);
      _outlen^:=a1;
      Move(PAnsiChar(rs)[0],_outbuf[0],a1);
    finally
        FreeMem(r,a1);
    end;
end;
procedure cBase64Encode(_in:PAnsiChar;_inlen:Integer;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;
var
r:AnsiString;
begin
    Base64Encode(_in,_inlen,r);
    StrCopy(_outbuf,PAnsiChar(r));
    _outlen^:=Length(r);
end;
procedure cEncodeString(_in:PAnsiChar;_key:PAnsiChar;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;
var
r:string;
begin
    r:=EncodeString(StrPas(_in),StrPas(_key));
    StrCopy(_outbuf,PAnsiChar(r));
    _outlen^:=Length(r);
end;
procedure cDecodeString(_in:PAnsiChar;_key:PAnsiChar;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;
var
r:string;
begin
    r:=DecodeString(string(StrPas(_in)),string(StrPas(_key)));
    StrCopy(_outbuf,PAnsiChar(r));
    _outlen^:=Length(r);
end;
procedure cDecodeStream(_instm:pbyte;_inlen:Integer;_outstm:pbyte;var _outlen:integer;_key:PAnsiChar);stdcall;
var
key:string;
instm:TMemoryStream;
outstm:TMemoryStream;
begin

    instm:=TMemoryStream.create;
    outstm:=TMemoryStream.Create;
    try
       instm.WriteBuffer(_instm^,_inlen);
       instm.Position:=0;
       DecodeStream(instm,outstm,stringof(bytesof(_key)));
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
procedure cEncodeStream(_instm:pbyte;_inlen:Integer;_outstm:pbyte;var _outlen:integer;_key:PAnsiChar);stdcall;
var
instm:TMemoryStream;
outstm:TMemoryStream;
begin
//    instm:=TMemoryStream.create;
//    outstm:=TMemoryStream.Create;
//    try
//       instm.WriteBuffer(_instm^,_inlen);
//       instm.Position:=0;
//       EncodeStream(instm,outstm,_key);
//       outstm.Seek(0, soFromBeginning);
//       if _outstm<>nil then
//       begin
//           Move(outstm.Memory^, _outstm^, outstm.Size);
//       end;
//       _outlen := outstm.Size;
//    finally
//      instm.Free;
//      outstm.Free;
//    end;
end;


exports
cEncodeString,
cDecodeString,
cBase64Decode,
cBase64Encode,
cEncodeStream,
cDecodeStream;


end.
