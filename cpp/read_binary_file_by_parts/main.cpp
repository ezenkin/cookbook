#include <iostream>
#include <fstream>


using namespace std;

int main() {
    int readed_bytes = 0, buf_size = 4096;
    char buf[buf_size];
    ifstream f("some.file", ios::in|ios::binary|ios::ate);
    if(!f) {
        cerr << "Failed to open file" << endl;
        return -1;
    }

    int left_bytes = f.tellg();
    cout << "Total bytes to read: " << left_bytes << endl;
    f.seekg(0, ios::beg);

    while (left_bytes > buf_size) {
        f.read(buf, buf_size);
        left_bytes-=buf_size;
        readed_bytes+=buf_size;
        cout << "Read " << buf_size << " bytes" << "\tReaded bytes: " << readed_bytes << "\tUnreaded bytes: " << left_bytes << endl;
    }

    if(left_bytes > 0) {
        f.read(buf, left_bytes);
        readed_bytes += left_bytes;
        cout << "Read " << left_bytes << " bytes" << "\tReaded bytes: " << readed_bytes << endl;
    }

    cout << "Total bytes readed: " << readed_bytes << endl;
    f.close();
    return 0;
}
