#include <stdio.h>
#pragma once

typedef int fcb(int);
void lol(fcb *cb, int a);

int fcbx(int a);


typedef void A;
typedef int B;

void FA(A* a);
void FB(B* b);
