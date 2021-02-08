import subprocess, sys

def generate_parser():
    subprocess.run("antlr4 -Dlanguage=Go -o parser Turnkey.g4", shell=True)


if __name__ == "__main__":
    args = sys.argv

    if args[1] == "gen_parser":
        print("Generating parser...")
        generate_parser()
        print("Generated.")
