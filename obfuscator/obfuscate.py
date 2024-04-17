import sys
import rjsmin

def obfuscate_js(input_path, output_path):
    try:
        with open(input_path, 'r') as file:
            js_code = file.read()
        
        # Using rjsmin to minify the JavaScript
        minified_js = rjsmin.jsmin(js_code)
        
        with open(output_path, 'w') as file:
            file.write(minified_js)
        
        print("Minification successful!")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: python obfuscate_js.py <input_file> <output_file>")
        sys.exit(1)
    
    input_file, output_file = sys.argv[1], sys.argv[2]
    obfuscate_js(input_file, output_file)
    print("Obfuscation successful!")