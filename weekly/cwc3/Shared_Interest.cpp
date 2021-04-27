int maxShared(int friends_nodes, vector<int> friends_from, vector<int> friends_to, vector<int> friends_weight)
{
  map<pair<int,int>,int> group;
  int n = friends_from.size();
  for(int i =0; i<n; i++)
  {
    int x = friends_from[i];
    int y = friends_to[i];
    group[make_pair(x,y)]++;
  }
  int max_count = INT_MIN;
  int result = INT_MIN;

  for(auto i = group.begin(); i!=group.end(); i++)
  {
    pair<int,int> p = i->first;
    int val_1 = i->second;
    int val_2 =0;
    int x = p.first;
    int y = p.second;
    if(group.find({y,x})!=group.end())
    {
      val_2 = group[{y,x}];
    }
    int count = val_1 + val_2;
    if(count >= max_count){
      max_count = count;
      result = max(result, (x+y));
    }
  }
  return result;
}

