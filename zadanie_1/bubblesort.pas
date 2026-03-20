program BubbleSort;

type
	TNumArray = array[1..50] of integer;

var
	RandNumArray : TNumArray;
	ItemCount : integer;

procedure Assert(Condition : boolean; TestName : string);
begin
	if Condition then
		WriteLn('[PASS] ', TestName)
	else
		WriteLn('[FAIL] ', TestName);
end;

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

// Unit tests
procedure RunUnitTests();
var
	TestArr : TNumArray;
	i : integer;
	isSuccess : boolean;
begin
	// TEST 1
	PopulateArray(TestArr, 10, 20, 50);
	isSuccess := true;
	for i := 1 to 50 do
		if (TestArr[i] < 10) or (TestArr[i] > 20) then isSuccess := false;
	Assert(isSuccess, 'Test 1 (Populating Array): Numbers are generated withing given range');
		
	// TEST 2
	for i := 1 to 50 do TestArr[i] := -1;
	PopulateArray(TestArr, 0, 1, 10);
	isSuccess := true;
	for i := 11 to 50 do
		if TestArr[i] <> -1 then isSuccess := false;
	Assert(isSuccess, 'Test 2 (Populating Array): Populating a section of array does not modify the rest');

	// TEST 3
	TestArr[1] := 6; TestArr[2] := 1; TestArr[3] := 2; TestArr[4] := 9; TestArr[5] := 3;
	SortArray(TestArr, 5);
	isSuccess := (TestArr[1]=1) and (TestArr[2]=2) and (TestArr[3]=3) and (TestArr[4]=6) and (TestArr[5]=9);
	Assert(isSuccess, 'Test 3 (Sorting): A subsection of an unsorted array is sorted correctly');

	// TEST 4
	TestArr[1] := 1; TestArr[2] := 2; TestArr[3] := 3;
	SortArray(TestArr, 3);
	isSuccess := (TestArr[1]=1) and (TestArr[2]=2) and (TestArr[3]=3);
	Assert(isSuccess, 'Test 4 (Sorting): A subsection of a sorted array remains sorted');

	// TEST 5
	for i := 1 to 50 do TestArr[i] := 0;
	TestArr[1] := 100; TestArr[25] := 99; TestArr[50] := 98;
	SortArray(TestArr, 50);
	isSuccess := (TestArr[1]=0) and (TestArr[47]=0) and (TestArr[48]=98) and (TestArr[49]=99) and (TestArr[50]=100);
	Assert(isSuccess, 'Test 5 (Sorting): Entire array is sorted correctly');

	WriteLn;
end;

begin
	ItemCount := 50;
	Randomize;

	RunUnitTests();

	PopulateArray(RandNumArray, 10, 30, ItemCount);
	PrintArray(RandNumArray);
	WriteLn;
	SortArray(RandNumArray, ItemCount);
	PrintArray(RandNumArray);
end.

