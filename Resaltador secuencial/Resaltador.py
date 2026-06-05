import re

PalabrasReservadas = ["if", "else", "while", "match", "switch", "case", "return", "lambda"]
OperadoresAritmeticos = ["+", "-", "*", "/"]
OperadoresComparacion = ["==", "<", ">", "<=", ">=", "!=", "/="]
OperadoresLogicos = ["and", "or", "not", "&&", "||", "!"]
Booleanos = ["true", "false", "True", "False"]
Nulos = ["None", "nullptr"]


def DetectarLenguaje(archivo):
    if archivo.endswith(".py"):
        return "python"
    elif archivo.endswith(".cpp") or archivo.endswith(".h"):
        return "cpp"
    elif archivo.endswith(".erl"):
        return "erlang"
    else:
        return "python"


def ConstruirPatron(lenguaje):
    keywords = {
        "python": r"(?P<KEYWORD>\b(?:if|else|while|match|case|return|lambda)\b)",
        "cpp":    r"(?P<KEYWORD>\b(?:if|else|while|switch|case|return)\b)",
        "erlang": r"(?P<KEYWORD>\b(?:if|case|of|end|fun)\b)",
    }
    operadores = {
        "python": r"(?P<LOGICO>\band\b|\bor\b|\bnot\b)",
        "cpp":    r"(?P<LOGICO>&&|\|\||!)",
        "erlang": r"(?P<LOGICO>\band\b|\bor\b|\bnot\b)",
    }
    comparacion = {
        "python": r"(?P<COMP>==|!=|<=|>=|<|>)",
        "cpp":    r"(?P<COMP>==|!=|<=|>=|<|>)",
        "erlang": r"(?P<COMP>==|/=|=<|>=|<|>)",
    }

    return "|".join([
        r"(?P<COMENTARIO>#.*|//.*|%.*)",
        r'(?P<STRING>"[^"]*"|\'[^\']*\')',
        r"(?P<BOOLEAN>true|false|True|False)",
        r"(?P<NULO>None|nullptr)",
        r"(?P<NUM>[0-9]+(?:\.[0-9]+)?)",
        keywords[lenguaje],
        comparacion[lenguaje],
        operadores[lenguaje],
        r"(?P<ARITMETICO>[+\-*/])",
        r"(?P<ID>[a-zA-Z_][a-zA-Z0-9_]*)",
    ])


def highlight(codigo, lenguaje):
    patron = ConstruirPatron(lenguaje)

    for m in re.finditer(patron, codigo):
        token = m.group()
        tipo  = m.lastgroup

        if tipo == "KEYWORD":
            print(f"\033[92m{token}\033[0m", end=" ")   # verde

        elif tipo == "LOGICO":
            print(f"\033[93m{token}\033[0m", end=" ")   # amarillo

        elif tipo in ("ARITMETICO", "COMP"):
            print(f"\033[91m{token}\033[0m", end=" ")   # rojo

        elif tipo == "COMENTARIO":
            print(f"\033[90m{token}\033[0m", end=" ")   # gris

        elif tipo == "STRING":
            print(f"\033[95m{token}\033[0m", end=" ")   # magenta

        elif tipo == "BOOLEAN":
            print(f"\033[96m{token}\033[0m", end=" ")   # cyan

        elif tipo == "NULO":
            print(f"\033[94m{token}\033[0m", end=" ")   # azul

        elif tipo == "NUM":
            print(f"\033[93m{token}\033[0m", end=" ")   # amarillo
            
        else:
            print(token, end=" ")
    print()


archivos = ["codigo.py", "codigo.cpp", "codigo.erl"]

for archivo in archivos:
    try:
        lenguaje = DetectarLenguaje(archivo)
        with open(archivo) as f:
            contenido = f.read()
        print(f"\n=== {archivo} ({lenguaje.upper()}) ===")
        highlight(contenido, lenguaje)
    except FileNotFoundError:
        print(f"✗ {archivo} no encontrado")