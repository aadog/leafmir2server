program Projecttest;

{$APPTYPE CONSOLE}

uses
  SysUtils,
  Unittest in 'Unittest.pas';



begin

  try
    { TODO -oUser -cConsole Main : Insert code here }

//    Unittest.testcBase64Encode;
    Unittest.testcBase64Decode;
    Readln;
  except
    on E: Exception do
    Writeln(E.ClassName, ': ', E.Message);
  end;
  Readln;
end.
