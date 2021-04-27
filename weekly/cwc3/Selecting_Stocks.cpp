int selectStock (int saving, int currentValue[], int futureValue[])
{
  int n = sizeof(futureValue)/sizeof(futureValue[0]);
  int profit_dp[n+1][saving+1];

  for(int i = 0; i<= n; i++)
  {
    for(int w = 0; w<= saving; w++)
    {
      if(i==0 || w==0)
         profit_dp[i][w]=0;
      else if (currentValue[i-1] <= w)
         profit_dp[i][w] = max((futureValue[i-1]-currentValue[i-1])+profit_dp[i-1][w-currentValue[i-1]], profit_dp[i-1][w]);
      else
         profit_dp[i][w] = profit_dp[i-1][w];
    }
  }

  return profit_dp[n][saving];
}
