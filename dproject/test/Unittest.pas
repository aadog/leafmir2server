unit Unittest;

interface

procedure cEncodeString(_in:PAnsiChar;_key:PAnsiChar;outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cEncodeString';
procedure cDecodeString(_in:PAnsiChar;_key:PAnsiChar;outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cDecodeString';
procedure cBase64Encode(_in:PAnsiChar;_inlen:Integer;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cBase64Encode';
procedure cBase64Decode(_in:PAnsiChar;_outbuf:PAnsiChar;_outlen:pInteger);stdcall;external '../Project.dll' name 'cBase64Decode';
procedure testcDecodeString();
procedure testcBase64Encode();
procedure testcBase64Decode();
implementation
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

end.
