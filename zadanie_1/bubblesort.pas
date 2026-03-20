program BubbleSort;

type
	TNumArray = array[1..50] of integer;

var
	RandNumArray : TNumArray;

procedure PrintArray(arr : TNumArray);
var
	idx : integer;
begin
	for idx := 1 to 50 do
	begin
		Write(arr[idx], ' ');
	end;
	WriteLn;
end;

procedure PopulateArray(var arr : TNumArray);
var
	idx : integer;
begin
	for idx := 1 to 50 do
	begin
		arr[idx] := Random(101);
	end;
end;

begin
	Randomize;
	PopulateArray(RandNumArray);
	PrintArray(RandNumArray);
end.
