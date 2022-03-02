#define _USE_MATH_DEFINES
#define midPointFile "midpoint_results.csv"
#define runge_kuttaFile "runge_kutta_results.csv"
#define iter_numb 30

#include <iostream>
#include <fstream>
#include <cmath>
#include <vector>

using namespace std;


const double deltaT = 0.125f;
const double gVal = -10;
const double mass = 1;
const double radius = 1;

double convertToRadians(double angle){
    return angle * M_PI / 180;
}

void writeResult(ofstream &ofs, double x, double y, double w){
    double potential = std::abs(mass * gVal * (radius + y));
    double kinetic = mass * (w * w * radius * radius) / 2;
    double total = potential + kinetic;

    ofs << potential << ", " << kinetic << ", " << total << endl;
}


pair<double, double> calculateDerivative(vector<double> &awe, pair<double, double> last_Ks, double dt){
    pair<double, double> result{};

    double a2 = 0;
    double w2 = 0;
    double e2 = 0;

    w2 = awe[1] + last_Ks.second * dt;
    a2 = awe[0] + last_Ks.first * dt;
    e2 = gVal / radius * sin(a2);

    result.first = w2;
    result.second = e2;

    return result;
}

void Runge_Kutta(){
    std::ofstream result_file(runge_kuttaFile);
    vector<double> init{convertToRadians(45), 0, 0};

    writeResult(result_file, radius * cos(init[0] - M_PI_2), radius * sin(init[0] - M_PI_2), init[1]);

    int i = 0;
    while (i < iter_numb){
        pair<double, double> dxdy{};
        vector<pair<double, double>> result{};

        dxdy = calculateDerivative(init, dxdy, 0);
        result.push_back(dxdy);
        dxdy = calculateDerivative(init, dxdy, deltaT / 2);
        result.push_back(dxdy);
        dxdy = calculateDerivative(init, dxdy, deltaT / 2);
        result.push_back(dxdy);
        dxdy = calculateDerivative(init, dxdy, deltaT);
        result.push_back(dxdy);
            
        init[0] += ((result[0].first + 2 * result[1].first + 2 * result[2].first + result[3].first) / 6) * deltaT;
        init[1] += ((result[0].second + 2 * result[1].second + 2 * result[2].second + result[3].second) / 6) * deltaT;
  
        writeResult(result_file, radius * cos(init[0] - M_PI_2), radius * sin(init[0] - M_PI_2), init[1]);
        i++;
    }
}

void mid_point(){
    std::ofstream result_file(midPointFile);

    double a = convertToRadians(45);
    double a2 = 0;
    double da = 0;
    double w = 0;
    double w2 = 0;
    double dw = 0;
    double e = 0;
    double e2 = 0;

    writeResult(result_file, radius * std::cos(a - M_PI_2), radius * std::sin(a - M_PI_2), w);

    int i = 0;
    for (i = 0 ; i < iter_numb; i++){
        e = gVal / radius * sin(a);
        w2 = w + e * deltaT / 2;
        da = w2 * deltaT;

        a2 = a + w * deltaT / 2;
        e2 = gVal / radius * sin(a2);
        dw = e2 * deltaT;

        a += da;
        w += dw;

        writeResult(result_file, radius * cos(a - M_PI_2), radius * sin(a - M_PI_2), w);
    }
}

int main(int argc, char const *argv[]){
    Runge_Kutta();
    mid_point();
    return 0;
}