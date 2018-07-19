#include <signal.h>  // signal()
#include <unistd.h> // _exit()

/* catch ctrl-c and destruct */
void stop(int signo) {
    _exit(0);
}

int main() {
    signal(SIGINT, stop);
}
