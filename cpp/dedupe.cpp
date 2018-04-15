/* c++ (std lib 11+) program to dedupe a sequence of ints with allowance for varying
thresholds of duplication */
#include <iostream>
#include <vector>

using namespace std;

// print out contents of vector<int> and add newline
void rep_i_vec(std::vector<int> *i_vec_p, int n) {
    int i = 0;
    std::vector<int> i_vec = *(i_vec_p);
    while (i < n) {
        cout << i_vec[i] << " ";
        i++;
    }
    cout << "\n";
}

// removes duplicates. considers duplication threshold
int remove_duplicates(std::vector<int> *i_vec, int n, int m_dupes)
{
    if (n==0 || n==1)
        return n;
 
    // j marks the extent of our distinct(ish) subarray
    int j = 0;

    /* create a reference shortcut to i_vec
    doing (*i_vec) again and again was ugly*/
    std::vector<int>& i_vec_ref = (*i_vec); 

    for (int i=0; i < n-1; i++) {
        // check if we have moved on in distinction
        if (i_vec_ref[i] != i_vec_ref[i+m_dupes])
            /* set the element at j to the element at i
            then increment j
            this is the default action (i.e. m_dupes = 1) */ 
            i_vec_ref[j++] = i_vec_ref[i];

            // move the good stuff over
            for (int b=1; b < m_dupes; b++) {
                i_vec_ref[j + b - 1] = i_vec_ref[i + b];
            }
    }
 
    // move the last element over
    i_vec_ref[j++] = i_vec_ref[n-1];
    return j;
}
 
int main()
{

    std::vector<int> i_vec = std::vector<int>();
    int count;
    int current;
    int max_dupes;

    cin >> count >> max_dupes;
    for(int i=0; i < count; i++) {
        cin >> current;
        i_vec.push_back(current);
    }
    int n = i_vec.size();
 
    // remove_duplicates returns size of de-duped subarray (vector)
    n = remove_duplicates(&i_vec, n, max_dupes);
 
    // Print updated i_vec
    rep_i_vec(&i_vec, n);
 
    return 0;
}