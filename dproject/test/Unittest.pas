unit Unittest;

interface
uses
Classes,
SysUtils,
EncdDecd;
type PTMemoryStream=^TMemoryStream;
procedure cEncodeString(_in:PAnsiChar;_key:PAnsiChar;outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cEncodeString';
procedure cDecodeString(_in:PAnsiChar;_key:PAnsiChar;outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cDecodeString';
procedure cBase64Encode(_in:PAnsiChar;_inlen:Integer;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cBase64Encode';
procedure cBase64Decode(_in:PAnsiChar;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cBase64Decode';
procedure cDecodeStream(_instm:pbyte;_inlen:Integer;_outstm:pbyte;var _outlen:integer;_key:PAnsiChar);stdcall;external '../Project.dll' name 'cDecodeStream';
procedure cEncodeStream(_instm:pbyte;_inlen:Integer;_outstm:pbyte;var _outlen:integer;_key:PAnsiChar);external '../Project.dll' name 'cEncodeStream';


procedure testcDecodeString();
procedure testcBase64Encode();
procedure testcBase64Decode();
procedure testbuf(_buf:PAnsiString);
implementation
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
//test
procedure testcDecodeString();
var
inbuf,key:AnsiString;
blen:Integer;
outbuf:AnsiString;
begin
    inbuf:='nimhJaix1yjuBzW0XWdKBZrLY5sL58Hsngtz2ud0uCJOtV2aQhFf6PRlUYNV+Q7OBSW2dy2vaxYZngIl7xekDAXdSUghiaxC90xJFcTR2uU3DT8S';
    key:='#$Ggy%(*^45fghj@@#sqw[]KHG%^&UHBR#$ty';
    SetLength(outbuf,255);
    cDecodeString(PAnsiChar(inbuf),PAnsiChar(key),PAnsiChar(outbuf),@blen);
    SetLength(outbuf,blen);
    Writeln(outbuf);
end;
procedure testcBase64Encode();
var
inbuf:string;
blen:Integer;
outbuf:AnsiString;
begin
    inbuf:='aaaa';
    SetLength(outbuf,255);
    cBase64Encode(PAnsiChar(inbuf),length(inbuf),PAnsiChar(outbuf),@blen);
    SetLength(outbuf,blen);
    Writeln(outbuf);
end;

procedure testcBase64Decode();
var
inbuf:AnsiString;
blen:Integer;
outl1:Integer;
outbuf:AnsiString;
begin
    inbuf:='>~CvKhKvKhK<<<<FWRCwKxCwKu<<<<bfa~PkWVLk<<<<<<'+'<<<<<<<<<<<<]~LR<vKBCrKROzLRCvKO<<<<<<<<<<<<<<<<<<<<Tfa~PkWO<<<<<<<<<<<<<<<<<<<<ffa~Pka~DiXe<<<<<Ka~DiXlLfXEXxXEXf<<<<<<<<<<<Ja~DiXkDxXlLfXEW<AhC&MRitKBCtKBC<<<<<<<<<'+'<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<';
    SetLength(outbuf,1080);
    FillChar(PAnsiChar(outbuf)^,1080,#0);
    cBase64Decode(PAnsiChar(inbuf),PAnsiChar(outbuf),@blen);
    Writeln(PAnsiChar(outbuf));
    Writeln(PAnsiChar(outbuf)[9]);
    Writeln(outbuf);
end;

procedure testbuf(_buf:PAnsiString);
var
inbuf:TBytes;
outbuf:TBytes;
outlen:Integer;
begin
    inbuf:=BytesOf('aaaaaaa');
    cEncodeStream(pbyte(@inbuf[0]),Length(inbuf),pbyte(@outbuf[0]),outlen,pansichar('KeyString'));
end;


end.
