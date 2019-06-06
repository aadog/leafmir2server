unit Unitcrypto;

interface
uses
uTPLb_Codec,uTPLb_CryptographicLibrary,uTPLb_Constants,SysUtils;
procedure Base64Encode(const ABuffer: PAnsiChar; ADataLen: Integer; out Result: AnsiString);
procedure Base64Decode(const ABase64Input: AnsiString; out ABuffer: PAnsiChar; out ADataLen, ABufferLen: Integer);
function EncodeString(const Source:AnsiString; Key: String): AnsiString;
function DecodeString(const Source: AnsiString; const Key: String): String;
implementation
const
  Base64_CodeTable: array [0..64] of AnsiChar = (
    '<', '>', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
    'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd',
    'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
    'u', 'v', 'w', 'x', 'y', 'z', '~', '[', ']', '&', '^', '|', '?', '@', '{', '}', '=');
  Base64_DecodeTable: array [0..127] of Byte = (
    255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
    255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
    255, 255, 255, 255, 255, 255, 57, 255, 255, 255, 255, 255, 255, 255, 255, 255,
    255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 0, 64, 1, 60, 61,
    2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
    24, 25, 26, 27, 55, 255, 56, 58, 255, 255, 28, 29, 30, 31, 32, 33, 34, 35, 36,
    37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 62, 59, 63,
    54, 255);
procedure Base64Encode(const ABuffer: PAnsiChar; ADataLen: Integer; out Result: AnsiString);
var
  Count: Integer;
  I: Integer;
begin
  I := (ADataLen + 2) div 3 * 4;
  SetLength(Result, I);
  I := 0;
  Count := 0;
  while Count < ADataLen do
  begin
    Inc(I);
    Result[I] := Base64_CodeTable[(Byte(ABuffer[Count]) and $FC) shr 2];
    if (Count + 1) < ADataLen then
    begin
      Inc(I);
      Result[I] := Base64_CodeTable[((Byte(ABuffer[Count]) and $03) shl 4) + ((Byte(ABuffer[Count + 1]) and $F0) shr 4)];
      if (Count + 2) < ADataLen then
      begin
        Inc(I);
        Result[I] := Base64_CodeTable[((Byte(ABuffer[Count + 1]) and $0F) shl 2) + ((Byte(ABuffer[Count + 2]) and $C0) shr 6)];
        Inc(I);
        Result[I] := Base64_CodeTable[(Byte(ABuffer[Count + 2]) and $3F)];
      end
      else
      begin
        Inc(I);
        Result[I] := Base64_CodeTable[(Byte(ABuffer[Count + 1]) and $0F) shl 2];
        Inc(I);
        Result[I] := '=';
      end
    end
    else
    begin
      Inc(I);
      Result[I] := Base64_CodeTable[(Byte(ABuffer[Count]) and $03) shl 4];
      Inc(I);
      Result[I] := '=';
      Inc(I);
      Result[I] := '=';
    end;
    Inc(Count, 3);
  end;
end;
procedure Base64Decode(const ABase64Input: AnsiString; out ABuffer: PAnsiChar; out ADataLen, ABufferLen: Integer);
var
  Count: Integer;
  Len: Integer;
  I: Integer;
  DataIn0: Byte;
  DataIn1: Byte;
  DataIn2: Byte;
  DataIn3: Byte;
begin
  ABufferLen := Length(ABase64Input) div 4 * 3 + 1;
  GetMem(ABuffer, ABufferLen);
  ABufferLen := ABufferLen;
  if ABuffer <> nil then
  begin
    Count := 1;
    Len := Length(ABase64Input);
    I := 0;
    while Count <= Len do
    begin
      if Byte(ABase64Input[Count]) in [13, 10] then
        Inc(Count)
      else
      begin
        DataIn0 := Base64_DecodeTable[Byte(ABase64Input[Count])];
        DataIn1 := Base64_DecodeTable[Byte(ABase64Input[Count + 1])];
        DataIn2 := Base64_DecodeTable[Byte(ABase64Input[Count + 2])];
        DataIn3 := Base64_DecodeTable[Byte(ABase64Input[Count + 3])];
        ABuffer[I] := AnsiChar(((DataIn0 and $3F) shl 2) + ((DataIn1 and $30) shr 4));
        Inc(I);
        if DataIn2 <> $40 then
        begin
          ABuffer[I] := AnsiChar(((DataIn1 and $0F) shl 4) + ((DataIn2 and $3C) shr 2));
          Inc(I);
          if DataIn3 <> $40 then
          begin
            ABuffer[I] := AnsiChar(((DataIn2 and $03) shl 6) + (DataIn3 and $3F));
            Inc(I);
          end;
        end;
        Count := Count + 4;
      end;
    end;
    ADataLen := I;
    ABuffer[I] := #0;
  end;
end;
function EncodeString(const Source:AnsiString; Key: String): AnsiString;
var
  ACode: TCodeC;
  ALibrary: TCryptographicLibrary;
begin
  ACode     :=  TCodeC.Create(nil);
  ALibrary  :=  TCryptographicLibrary.Create(nil);
  try
    ACode.CryptoLibrary  :=  ALibrary;
    ACode.AsymetricKeySizeInBits := 1024;
    ACode.StreamCipherId := BlockCipher_ProgId;
    ACode.BlockCipherId  :=  Blowfish_ProgId;
    ACode.ChainModeId  :=  CBC_ProgId;
    ACode.Password :=  Key;
    ACode.EncryptString(Source, Result);
  finally
    FreeAndNil(ACode);
    FreeAndNil(ALibrary);
  end;
end;
function DecodeString(const Source: AnsiString; const Key: String): String;
var
  ACode: TCodeC;
  ALibrary: TCryptographicLibrary;
begin
  ACode     :=  TCodeC.Create(nil);
  ALibrary  :=  TCryptographicLibrary.Create(nil);
  try
    ACode.CryptoLibrary  :=  ALibrary;
    ACode.AsymetricKeySizeInBits := 1024;
    ACode.StreamCipherId := BlockCipher_ProgId;
    ACode.BlockCipherId  :=  Blowfish_ProgId;
    ACode.ChainModeId  :=  CBC_ProgId;
    ACode.Password :=  Key;
    ACode.DecryptString(Result, Source);
  finally
    FreeAndNil(ACode);
    FreeAndNil(ALibrary);
  end;
end;
end.
