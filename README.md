# Cyclomatix
<pre style="white-space: pre;" align="center">
 ▄▄▄▄▄▄▄▄▄▄▄  ▄         ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄            ▄▄▄▄▄▄▄▄▄▄▄ 
▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░▌          ▐░░░░░░░░░░░▌
▐░█▀▀▀▀▀▀▀▀▀ ▐░▌       ▐░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░▌          ▐░█▀▀▀▀▀▀▀█░▌
▐░▌          ▐░▌       ▐░▌▐░▌          ▐░▌          ▐░▌       ▐░▌
▐░▌          ▐░█▄▄▄▄▄▄▄█░▌▐░▌          ▐░▌          ▐░▌       ▐░▌
▐░▌          ▐░░░░░░░░░░░▌▐░▌          ▐░▌          ▐░▌       ▐░▌
▐░▌           ▀▀▀▀█░█▀▀▀▀ ▐░▌          ▐░▌          ▐░▌       ▐░▌
▐░▌               ▐░▌     ▐░▌          ▐░▌          ▐░▌       ▐░▌
▐░█▄▄▄▄▄▄▄▄▄      ▐░▌     ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌
▐░░░░░░░░░░░▌     ▐░▌     ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
 ▀▀▀▀▀▀▀▀▀▀▀       ▀       ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀ 

 ▄▄       ▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄       ▄   
▐░░▌     ▐░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌     ▐░▌  
▐░▌░▌   ▐░▐░▌▐░█▀▀▀▀▀▀▀█░▌ ▀▀▀▀█░█▀▀▀▀  ▀▀▀▀█░█▀▀▀▀  ▐░▌   ▐░▌   
▐░▌▐░▌ ▐░▌▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌       ▐░▌ ▐░▌    
▐░▌ ▐░▐░▌ ▐░▌▐░█▄▄▄▄▄▄▄█░▌     ▐░▌          ▐░▌        ▐░▐░▌     
▐░▌  ▐░▌  ▐░▌▐░░░░░░░░░░░▌     ▐░▌          ▐░▌         ▐░▌      
▐░▌   ▀   ▐░▌▐░█▀▀▀▀▀▀▀█░▌     ▐░▌          ▐░▌        ▐░▌░▌     
▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌       ▐░▌ ▐░▌    
▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌      ▄▄▄▄█░█▄▄▄▄  ▐░▌   ▐░▌   
▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌     ▐░░░░░░░░░░░▌▐░▌     ▐░▌  
 ▀         ▀  ▀         ▀       ▀       ▀▀▀▀▀▀▀▀▀▀▀  ▀       ▀   
</pre>
<pre align= "center">
A Go static analysis tool to generate control flow graphs and compute cyclomatic complexity
</pre>

## Features

### Cyclomatic Complexity Computation

Cyclomatix computes the cyclomatic complexity of each and every function found in the input files given by the user to the tool

### Control Flow Graph Generation

Cyclomatix traverses all of the functions found in the files inputted by the users to generate their control flow graphs then outputs them in DOT files used by Graphviz. 

> [!WARNING]  
> To fully use the control flow graph generation feature, the user must install Graphviz in their machine.

## Installation guide

1. Download the latest release that corresponds with your system from the [releases page](https://github.com/Assifar-Karim/cyclomatix/releases).
2. Decompress the archive containing the binaries.
3. Install Graphviz on your system, if it's not already installed, by following the instructions found [here](https://graphviz.org/download/). 
4. Add the binary to your PATH environment variable
5. Enjoy 

## Getting started

After having installed cyclomatix on your system, you can follow the steps to get started on using the tool.

1. Pull the `.go` files from the `examples` directory in this repo.
2. Run the command `cyclo complexity -f examples` to get the cyclomatic complexity table of the functions in the files.
3. Run the command `cyclo cfg -f example -o target` to generate the control flow graph of each function that can be found on the example files.  