import java.utils.*;
import java.io.*;

class countingValleys{
    public static void main(String[] args){
        BufferedReader bf = new BufferedReader(new InputStreamReader(System.in));
        System.out.println("How many steps?");
        int steps=bf.readLine();
        System.out.println("String for path (U=uphill, D=downhill)");
        String path = bf.readLine();
        System.out.println(countingValleys(steps, path));

    }

    public static int countingValleys(int steps, String path) {
        int altitudeP = 0;
        int altitudeC = 0;
        int valleys = 0;
        
        if(path.charAt(0)=='U'){
            altitudeP++;
            altitudeC++;
        }else{
            altitudeP--;
            altitudeC--;
        }

        for(int i=1; i<path.length(); i++){
            char step=path.charAt(i);
            if(step=='U'){
                altitudeC++;
                if(altitudeP<0 && altitudeC==0)
                    valleys++;
            }else{
                altitudeC--;
            }    
            altitudeP=altitudeC;
        }
        
        return valleys;
    }
}