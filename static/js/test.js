
var allListings = document.querySelectorAll(".workoutlistings");
var workoutCount = document.querySelector("#workoutcount");
var tabLength = allListings.length;

workoutCount.innerText = "Total Workouts Completed: " + tabLength;

const allworkoutDates = [];
const uniqueDates =[];

allListings.forEach((workout) => {
    const regExp = /\(([^)]+)\)/g;
    const matches = [...workout.innerText.match(regExp)];
    const arr = matches[0].split("/");
    let month = arr[0];
    //console.log(arr);
    allworkoutDates.push(month.substring(1)+"/23");
    
});

const counts = [];
allworkoutDates.forEach((month) => {
    let count = 0;
    for (let i = 0; i < allworkoutDates.length; i++){
        if (month == allworkoutDates[i]) {
            count += 1;
        }
    }
    counts.push(count);
});

const months = [];
const monthcount = [];
var count = 0;
     
    var start = false;
     
    for (j = 0; j < allworkoutDates.length; j++) {
        for (k = 0; k < months.length; k++) {
            if ( allworkoutDates[j] == months[k] ) {
                start = true;
            }
        }
        count++;
        if (count == 1 && start == false) {
            months.push(allworkoutDates[j]);
        }
        start = false;
        count = 0;
    }

    var start = false;
     
    for (j = 0; j < counts.length; j++) {
        for (k = 0; k < monthcount.length; k++) {
            if ( counts[j] == monthcount[k] ) {
                start = true;
            }
        }
        count++;
        if (count == 1 && start == false) {
            monthcount.push(counts[j]);
        }
        start = false;
        count = 0;
    }

console.log(months);
console.log(monthcount);


const ctx = document.getElementById("myChart");

              new Chart(ctx, {
                type:"bar",
                data: {
                  labels: months,
                  datasets: [
                    {
                      label: "# Of Workouts Per Month",
                      data: monthcount,
                      borderWidth: 1,
                    },
                  ],
                },
                options: {
                  scales: {
                    y: {
                      beginAtZero: true,
                    },
                  },
                },
              });

