# include <bits/stdc++.h>

using namespace std;
const int a = 69069;
const int b = 5;
const long long m = (1ll << 32);
const int cnt_test = 1; // Change it if u wanna make tests
double Gf,Lf;
int main(){
    srand (time(NULL));
    long long x,length;
    printf("Please enter a initial number of the sequence\nIf you wanna random initial number enter -1 : ");
    scanf("%I64d",&x);
    if(x == -1){
        x = rand();
    }
    printf("\nYou can choose the length of the sequence or enter -1 to randomize it from 1 to 10000000\n");
    scanf("%I64d",&length);
    if(length == -1){
        length = rand() % 10000001;
    }
    printf("\nInitial values of [x=(a*x+c)%m] is a = %d, b = %d, m = %I64d, x - %I64d, length - %d\n\n",a,b,m,x,length);
    for(int test = 1; test <= cnt_test;test ++){
    
    vector<long long> nums;
    for(int i = 1;i <= length;i ++){
        x = (a*x+b) % m;
        nums.push_back(x);
        //printf("%d-th - %I64d\n",i,x);
    }
    sort(nums.begin(),nums.end());
    int duplicates = 0;
    int Great, Less;
    Great = Less = 0;
    for(int i = 0;i < nums.size();i ++){
        if(i - 1 >= 0 && nums[i - 1] == nums[i]) duplicates++;
        if(nums[i] >= m / 2ll) Great++;
        else Less++;
    }
    if(cnt_test != 1){
        printf("Test %d: length - %I64d\n",test,length);
        length = rand() % 10000001;
    }
    printf("In the sequence %d - duplicates and %d - unique numbers\n",duplicates,nums.size() - duplicates);
    printf("Where %lf%c numbers - Greater than m/2(%I64d) and %lf%c numbers - Less than m/2(%I64d)\n\n",Great * 100.0/nums.size(),'%',m/2ll,Less*100.0/nums.size(),'%',m/2ll);
    Gf += (Great * 100.0/nums.size());
    Lf += (Less * 100.0/nums.size());
    }
    if(cnt_test != 1){
        Gf /= cnt_test;
        Lf /= cnt_test;
        printf("Final probability %lf%c numbers - Greater than m/2 and %lf%c numbers - Less than m/2\n",Gf,'%',Lf,'%');
    }
}
//4294966769
//
