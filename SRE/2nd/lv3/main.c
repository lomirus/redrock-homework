#include <stdio.h>

char* find_char(char const* source, char const* chars) {
    if (source == NULL || chars == NULL)
        return NULL;
    if (*source == NULL || *chars == NULL)
        return NULL;
    for (char* ps = source;*ps;ps++)
        for (char* pc = chars;*pc;pc++)
            if (*ps == *pc)
                return ps;
	return NULL;
}


int main() {
    printf("%p", find_char("ABCDEF", "XRCQEF"));
    return 0;
}