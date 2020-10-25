#include <stdio.h>

int triPow(int n){
	return  n*n*n;
}

int main(){
	for(int i = 100;i < 1000;i++){
		int a = i / 100;
		int b = (i - a * 100)/10;
		int c = i % 10;
		if(i == triPow(a) + triPow(b) + triPow(c)){
			printf("%d\n",i);
		}
		
	}
	return 0;
} 
