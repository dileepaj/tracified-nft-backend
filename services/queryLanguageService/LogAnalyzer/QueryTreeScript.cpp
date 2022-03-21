// //
// // Created by Michelle on 12/31/2019.
// //

// #include <Node.h>
// #include "QueryTreeScript.h"
// #include <iostream>
// #include <fstream>
// #include <Node.h>
// #include "QueryExecuter.h"
// #include <algorithm>



// using namespace std;

// void QueryTreeScript::QueryNodeTree(Node* root){

//     ifstream queryFile ("../FlexibleComputerLanguage/QueryResult/Query.txt");
//     string query="";
//     string qline="";
//     while(getline(queryFile,qline))
//     {
//         query+=qline;
//         query+="\n";
//     }

//     string res=QueryExecuter::run(root,query);
//     cout << res;

//     ofstream resultfile;
//     resultfile.open("../FlexibleComputerLanguage/QueryResult/QueryResult.json");
//     resultfile << res;

// }
