# include <bits/stdc++.h>
using namespace std;
map<pair<int,int>,int>st;
bool check(int b[]){
    int c[12];
    for(int i = 0;i < 9;i ++){
        c[i] = b[i];
    }
    c[9] = b[0];
    c[10] = b[1];
    for(int i = 0;i <= 8; i ++){
        if((c[i] + c[i + 1] + c[i + 2]) % 3) return false;
    }
    return true;
}
int main(){
    int a[] = {6,47,16,26,31,67,63,74,120};
    sort(a,a+9);
    int cnt = 0;
     do{
        if (check(a)){
            cnt ++;
            st[make_pair(a[0],a[1])]++;
        }
    } while(next_permutation(a,a + 9));
    for(auto i: st){
        printf("Secret code start from %d,%d - total available codes is %d\n",i.first.first,i.first.second,i.second);
    }
    cout <<"Count of available arrays - " <<  cnt << endl;
    sort(a,a+9);
    do{
        if (check(a)){
            for(auto i : a){
                cout << i << ' ';
            }
            cout << endl;
        }
    } while(next_permutation(a,a + 9));
}