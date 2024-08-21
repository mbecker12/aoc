import sys
import sympy


if __name__ == "__main__":
    result1 = sys.argv[1]
    result2 = sys.argv[2]

    x = sympy.Symbol('x')
    expr1 = sympy.parsing.sympy_parser.parse_expr(result1)
    expr2 = sympy.parsing.sympy_parser.parse_expr(result2)
    if result2 == 'level1':
        solution = sympy.N(expr1 - expr2)
    else:
        solution = sympy.solvers.solve(expr1 - expr2)[0]
    print(f"{int(solution)}")