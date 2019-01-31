package main

import (
    "github.com/krolaw/xsd"
    "github.com/jbussdieker/golibxml"
    "fmt"
    "os"
    "unsafe"
)

func usage(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
	}
    fmt.Fprint(os.Stderr, "A xml schama validator tool\n\n")
    fmt.Fprint(os.Stderr, "usage: xmlValidator <xml_file|xml_dir> <xml_schema_file>\n\n")
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

    /*method one
    xsdb := make([]byte, 1024*1024*10)
    var xsdfile, err = os.Open(os.Args[2])
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }

    n, err := xsdfile.Read(xsdb)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
    defer xsdfile.Close()
    xsdb = xsdb[0:n:n]
    xsdSchema, err := xsd.ParseSchema(xsdb) */

    //method two 
    xsdSchema, err := xsd.ParseSchemaFile(os.Args[2])

    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }


    xmlbytes := make([]byte, 1024*1024*10)
    for _, f := range files {

        fi, err := os.Open(f)
        if err != nil {
            fmt.Println("open file: ", f, " failed", err.Error())
            continue
        }else {
            defer fi.Close()
        }

        nbytes, err := fi.Read(xmlbytes)
        if err != nil {
            fmt.Println("read file:", f, "failed")
            continue
        }else {
            XML := string(xmlbytes[0:nbytes:nbytes])
            doc := golibxml.ParseDoc(XML)
            if doc == nil {
                fmt.Println("Error parsing document:", f)
                continue
            }else {
                defer doc.Free()
                /* method one
                if err := xsdSchema.Validate(xsd.DocPtr(unsafe.Pointer(doc.Ptr))); err != nil {
                    fmt.Println(f, ":\n", err)
                }else {
                    fmt.Println(f, ":\n matched!\n", )
                }*/

                // method two
                if err := xsdSchema.ValidateS(xsd.DocPtr(unsafe.Pointer(doc.Ptr))); err != nil {
                    fmt.Println(f, ":\n", err)
                }else {
                    fmt.Println(f, ":\n matched!\n", )
                }

                /*//method three 
                if err := xsdSchema.ValidateFile(f); err != nil {
                    fmt.Println(f, ":\n", err)
                }else {
                    fmt.Println(f, ":\n matched!\n", )
                }*/
            }
        }

    }

}