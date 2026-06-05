-module(codigo).

test(X) ->
    case X of
        10 -> true;
        _ -> false
    end.

main() ->
    Y = 5,
    Z = 10,
    if Z == 10 and Y /= 0 ->
        io:format("Erlang funciona~n");
    true ->
        ok
    end.