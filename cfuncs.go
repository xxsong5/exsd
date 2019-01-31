package exsd

/*
#include <stdio.h>
#include<string.h>
#include <stdarg.h>
#include <libxml/tree.h>
#include <libxml/xmlerror.h>

// The gateway function
void xmlErrorFunc_cgo(void *ctx, const char * format, ...)
{
    void xmlErrorFunc(void *ctx, const char *);

    char buf[1024];
    va_list args;
    va_start(args, format);
    vsnprintf(buf, sizeof(buf), format, args);
    va_end(args);

    xmlErrorFunc(ctx, buf);
}

void xmlSErrorFunc_cgo(void *ctx, void *Serror)
{
    void xmlErrorFunc(void *ctx, const char *);

    xmlErrorPtr SerrorPtr = (xmlErrorPtr)Serror;

    xmlNodePtr node = (xmlNodePtr)SerrorPtr->node;

    char nodeInfo[512][512];
    int idx = 0;
    while(node && node->parent) {
        idx != 0 ? sprintf(nodeInfo[idx++], "%s", node->name): sprintf(nodeInfo[idx++], "%s (line:%d, errorCode:%d)", node->name, node->line, SerrorPtr->code);
        node = node->parent;
    }

    char nodeInfos[5120];
    char tag[]=", ";
    
    int iii = 0;
    while(--idx >= 0) {
        int len = strlen(nodeInfo[idx]) + strlen(tag);
        idx != 0 ? sprintf(nodeInfos+iii, "%s%s", nodeInfo[idx], tag) : sprintf(nodeInfos+iii, "%s", nodeInfo[idx]);
        iii += len;
    }

    char buf[1048];
    sprintf(buf, "\033[0;33m[%s]\033[0m \033[0;32m%s\033[0m\n", nodeInfos, SerrorPtr->message);

    xmlErrorFunc(ctx, buf);
}

*/
import "C"
