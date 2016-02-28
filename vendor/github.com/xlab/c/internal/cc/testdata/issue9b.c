typedef short JBLOCK[64];
typedef JBLOCK *JBLOCKROW;
typedef JBLOCKROW *JBLOCKARRAY;
typedef JBLOCKARRAY *JBLOCKIMAGE;

short (**ok)[64];
JBLOCKARRAY ok;

short ***bad[64];
JBLOCKARRAY bad;
