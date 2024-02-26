static check?
4_contextcheck/main.go:12:15: SA4009: argument ctx is overwritten before first use (staticcheck)
func someFunc(ctx context.Context) {
^
4_contextcheck/main.go:13:2: SA4009(related information): assignment to ctx (staticcheck)
ctx = context.Background()
