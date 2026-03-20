program BubbleSort;

type
	TNumArray = array[1..50] of integer;

var
	RandNumArray : TNumArray;
	ItemCount : integer;

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

procedure PopulateArray(var arr : TNumArray; minVal : integer; maxVal : integer; count : integer);
var
	idx : integer;
begin
	for idx := 1 to count do
	begin
		arr[idx] := Random(maxVal - minVal + 1) + minVal;
	end;
end;

procedure SortArray(var arr : TNumArray; count : integer);
var
	i, j, tmp : integer;
begin
	for i := 1 to count - 1 do
	begin
		for j := 1 to count - i do
		begin
			if arr[j] > arr[j + 1] then
			begin
				tmp := arr[j];
				arr[j] := arr[j + 1];
				arr[j + 1] := tmp;
			end;
		end;
	end;
end;

begin
	ItemCount := 50;
	Randomize;
	PopulateArray(RandNumArray, 10, 30, ItemCount);
	PrintArray(RandNumArray);
	WriteLn;
	SortArray(RandNumArray, ItemCount);
	PrintArray(RandNumArray);
end.
