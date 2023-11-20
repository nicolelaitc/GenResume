# GenResume

## Description
Let Latex take care of the formatting of your resume, and you just have to deal with texts. Especially if you love to let ChatGPT do all the tailoring for you :D. 

## Warnings
1. If you intend to tailor your resume with the help of AI tools, please only send the information listed in `sample.json` to avoid privacy issues. 
2. If you find the resume format of [ResuMake](https://resumake.io/) is good enough for you, you probably should use it instead.
3. The program (currently) cannot deal with `%` and `&` properly. Please replace all `%`/`&` characters with `//%`/`//&`.
4. Further customization requires LaTeX knowledge.

## Usage
You need to install [Go](https://golang.org/doc/install) and [pdflatex](https://www.scijournal.org/articles/latex-installation-guide) to make this work.
1. Download the code 
2. Update the `header.json`. Update `sample.json` or you can create a new json file to save your resume body. 
3. Open the terminal and enter

```bash
# Assuming you save the file folder in Downloads/ and use the `sample.json` for the resume body. 
cd Downloads/GenResume; go run *.go sample.json
```
If you want to edit the format, go to `template.tex` and `resume_config.cls`, like how you edit LaTeX. Please leave a comment (or find ChatGPT for assistance) if you have any difficulties formatting the template. 

## License
MIT

## Credits
I made this alone (with the help of ChatGPT). 
