package main

import (
    "github.com/xxsong5/exsd"
    "fmt"
    "os"
)

func usage(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
    fmt.Fprint(os.Stderr, "A xml schama validator tool\n\n")
    fmt.Fprint(os.Stderr, "usage: xmlValidator1 <xml_file|xml_dir> <xml_schema_file>\n\n")
    os.Exit(2)
}

func main()  {
    if l:=len(os.Args);l != 3 {
        usage("")
    }

	// check file type
	finfo1, err1 := os.Stat(os.Args[1])
	finfo2, err2 := os.Stat(os.Args[2])

	// Unable to find either file/directory
	if err1 != nil || err2 != nil {
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err1.Error())
		}
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err2.Error())
		}
		os.Exit(1)
	}
	if finfo2.IsDir() {
        usage("xml schama file can not be a direcory");
    }

    // read xml-file to files []string
    var files []string
    if l:=finfo1.IsDir(); l==false {
        files = append(files, os.Args[1])
    }else {
        dir, err := os.Open(os.Args[1])
        if err != nil {
            fmt.Fprintf(os.Stderr, "%s\n", err.Error())
            os.Exit(1)
        }
        names, err :=dir.Readdir(0)
        if err != nil {
            fmt.Fprintf(os.Stderr, "%s\n", err.Error())
            os.Exit(1)
        }
        rootdir := os.Args[1];
        if rootdir[len(rootdir)-1] != '/' {
            rootdir += "/"
        }
        for _, n := range names {
            if !n.IsDir() {
                files = append(files, rootdir+n.Name())
            }
        }
    }


    //load xsd file
    xsdSchema, err := exsd.ParseSchemaFile(os.Args[2])
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }

    //validate xml file
    for _, f := range files {

        if err := xsdSchema.ValidateFile(f); err != nil {
            fmt.Println(f, ":\n", err)
        }else {
            fmt.Println(f, ":\n matched!")
        }
    }

}