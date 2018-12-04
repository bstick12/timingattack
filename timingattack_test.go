package timingattack_test

import (
	"testing"

	"github.com/bstick12/timingattack"
)

func TestLoginConstantFailure(t *testing.T) {
	if timingattack.LoginConstant([]byte("BLAH")) {
		t.Errorf("Wanted login to fail")
	}
}

func TestLoginConstantSuccess(t *testing.T) {
	if !timingattack.LoginConstant([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")) {
		t.Errorf("Wanted login to be successful")
	}
}

func TestLoginFailure(t *testing.T) {
	if timingattack.Login("BLAH") {
		t.Errorf("Wanted login to fail")
	}
}

func TestLoginSuccess(t *testing.T) {
	if !timingattack.Login("ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		t.Errorf("Wanted login to be successful")
	}
}

func benchmarkLoginConstant(password string, b *testing.B) {
	bytes := []byte(password)
	for n := 0; n < b.N; n++ {
		timingattack.LoginConstant(bytes)
	}
}

func benchmarkLogin(password string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		timingattack.Login(password)
	}
}

func BenchmarkLoginKnownLength1(b *testing.B)  { benchmarkLogin("A.........................", b) }
func BenchmarkLoginKnownLength2(b *testing.B)  { benchmarkLogin("AB........................", b) }
func BenchmarkLoginKnownLength3(b *testing.B)  { benchmarkLogin("ABC.......................", b) }
func BenchmarkLoginKnownLength4(b *testing.B)  { benchmarkLogin("ABCD......................", b) }
func BenchmarkLoginKnownLength5(b *testing.B)  { benchmarkLogin("ABCDE.....................", b) }
func BenchmarkLoginKnownLength6(b *testing.B)  { benchmarkLogin("ABCDEF....................", b) }
func BenchmarkLoginKnownLength7(b *testing.B)  { benchmarkLogin("ABCDEFG...................", b) }
func BenchmarkLoginKnownLength8(b *testing.B)  { benchmarkLogin("ABCDEFGH..................", b) }
func BenchmarkLoginKnownLength9(b *testing.B)  { benchmarkLogin("ABCDEFGHI.................", b) }
func BenchmarkLoginKnownLength10(b *testing.B) { benchmarkLogin("ABCDEFGHIJ................", b) }
func BenchmarkLoginKnownLength11(b *testing.B) { benchmarkLogin("ABCDEFGHIJK...............", b) }
func BenchmarkLoginKnownLength12(b *testing.B) { benchmarkLogin("ABCDEFGHIJKL..............", b) }
func BenchmarkLoginKnownLength13(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLM.............", b) }
func BenchmarkLoginKnownLength14(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMN............", b) }
func BenchmarkLoginKnownLength15(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNO...........", b) }
func BenchmarkLoginKnownLength16(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOP..........", b) }
func BenchmarkLoginKnownLength17(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQ.........", b) }
func BenchmarkLoginKnownLength18(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQR........", b) }
func BenchmarkLoginKnownLength19(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRS.......", b) }
func BenchmarkLoginKnownLength20(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRST......", b) }
func BenchmarkLoginKnownLength21(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTU.....", b) }
func BenchmarkLoginKnownLength22(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUV....", b) }
func BenchmarkLoginKnownLength23(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVW...", b) }
func BenchmarkLoginKnownLength24(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVWX..", b) }
func BenchmarkLoginKnownLength25(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVWXY.", b) }
func BenchmarkLoginKnownLength26(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVWXYZ", b) }

func BenchmarkLogin1(b *testing.B)  { benchmarkLogin("A", b) }
func BenchmarkLogin2(b *testing.B)  { benchmarkLogin("AB", b) }
func BenchmarkLogin3(b *testing.B)  { benchmarkLogin("ABC", b) }
func BenchmarkLogin4(b *testing.B)  { benchmarkLogin("ABCD", b) }
func BenchmarkLogin5(b *testing.B)  { benchmarkLogin("ABCDE", b) }
func BenchmarkLogin6(b *testing.B)  { benchmarkLogin("ABCDEF", b) }
func BenchmarkLogin7(b *testing.B)  { benchmarkLogin("ABCDEFG", b) }
func BenchmarkLogin8(b *testing.B)  { benchmarkLogin("ABCDEFGH", b) }
func BenchmarkLogin9(b *testing.B)  { benchmarkLogin("ABCDEFGHI", b) }
func BenchmarkLogin10(b *testing.B) { benchmarkLogin("ABCDEFGHIJ", b) }
func BenchmarkLogin11(b *testing.B) { benchmarkLogin("ABCDEFGHIJK", b) }
func BenchmarkLogin12(b *testing.B) { benchmarkLogin("ABCDEFGHIJKL", b) }
func BenchmarkLogin13(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLM", b) }
func BenchmarkLogin14(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMN", b) }
func BenchmarkLogin15(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNO", b) }
func BenchmarkLogin16(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOP", b) }
func BenchmarkLogin17(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQ", b) }
func BenchmarkLogin18(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQR", b) }
func BenchmarkLogin19(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRS", b) }
func BenchmarkLogin20(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRST", b) }
func BenchmarkLogin21(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTU", b) }
func BenchmarkLogin22(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUV", b) }
func BenchmarkLogin23(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVW", b) }
func BenchmarkLogin24(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVWX", b) }
func BenchmarkLogin25(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVWXY", b) }
func BenchmarkLogin26(b *testing.B) { benchmarkLogin("ABCDEFGHIJKLMNOPQRSTUVWXYZ", b) }

func BenchmarkLoginConstant1(b *testing.B)  { benchmarkLoginConstant("A", b) }
func BenchmarkLoginConstant2(b *testing.B)  { benchmarkLoginConstant("AB", b) }
func BenchmarkLoginConstant3(b *testing.B)  { benchmarkLoginConstant("ABC", b) }
func BenchmarkLoginConstant4(b *testing.B)  { benchmarkLoginConstant("ABCD", b) }
func BenchmarkLoginConstant5(b *testing.B)  { benchmarkLoginConstant("ABCDE", b) }
func BenchmarkLoginConstant6(b *testing.B)  { benchmarkLoginConstant("ABCDEF", b) }
func BenchmarkLoginConstant7(b *testing.B)  { benchmarkLoginConstant("ABCDEFG", b) }
func BenchmarkLoginConstant8(b *testing.B)  { benchmarkLoginConstant("ABCDEFGH", b) }
func BenchmarkLoginConstant9(b *testing.B)  { benchmarkLoginConstant("ABCDEFGHI", b) }
func BenchmarkLoginConstant10(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJ", b) }
func BenchmarkLoginConstant11(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJK", b) }
func BenchmarkLoginConstant12(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKL", b) }
func BenchmarkLoginConstant13(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLM", b) }
func BenchmarkLoginConstant14(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMN", b) }
func BenchmarkLoginConstant15(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNO", b) }
func BenchmarkLoginConstant16(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOP", b) }
func BenchmarkLoginConstant17(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQ", b) }
func BenchmarkLoginConstant18(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQR", b) }
func BenchmarkLoginConstant19(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRS", b) }
func BenchmarkLoginConstant20(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRST", b) }
func BenchmarkLoginConstant21(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRSTU", b) }
func BenchmarkLoginConstant22(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRSTUV", b) }
func BenchmarkLoginConstant23(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRSTUVW", b) }
func BenchmarkLoginConstant24(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRSTUVWX", b) }
func BenchmarkLoginConstant25(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRSTUVWXY", b) }
func BenchmarkLoginConstant26(b *testing.B) { benchmarkLoginConstant("ABCDEFGHIJKLMNOPQRSTUVWXYZ", b) }
