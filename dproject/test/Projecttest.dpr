program Projecttest;

{$APPTYPE CONSOLE}

uses
  SysUtils,
  Unittest in 'Unittest.pas';


var
t1:AnsiString;
begin
  try
    { TODO -oUser -cConsole Main : Insert code here }

//    Unittest.testcBase64Encode;
    SetLength(t1,50);
    FillChar(PAnsiChar(t1)^,Length(t1),#0);

    Unittest.testbuf(@t1);
    Readln;
  except
    on E: Exception do
    Writeln(E.ClassName, ': ', E.Message);
  end;
  Readln;
end.
