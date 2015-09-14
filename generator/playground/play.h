#include <stdio.h>
#pragma once

typedef int fcb(int);
void lol(fcb *cb, int a);

int fcbx(int a);
