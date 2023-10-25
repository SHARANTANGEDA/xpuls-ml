// LineChart.tsx

import React, {useEffect, useRef, useState} from 'react';
import ReactEcharts from 'echarts-for-react';
import ReactECharts from "echarts-for-react";


interface LineChartProps {
    data: UsageDataResponse;
    title?: string;
    yAxisName: string;
}

const LineChart: React.FC<LineChartProps> = ({ data, title, yAxisName }) => {
    const [loaded, setLoaded] = useState(false);
    const echartInstance = useRef<ReactECharts | null>(null);

    // const transformedData = convertJsonToTimeSeries(data)
    useEffect(() => {

    }, []);

    let series: {dimension: string, values: number[]}[] = []
    const dimensions = Object.keys(data)
        .filter(key => key !== 'timestamp')

    dimensions.forEach((value, index, array) => {
            series.push({dimension: value, values: data[value]})
    })
    console.log(series)



    const getOption = () => {
        return {
            title: {
                text: title
            },
            tooltip: {
                trigger: 'axis'
            },
            legend: {
                data: dimensions
            },
            xAxis: {
                type: 'time',
                axisLabel: {
                    formatter: function (value: number, index: number) {
                        // 'value' here is the timestamp in milliseconds since the Unix epoch
                        const date = new Date(value * 1000);

                        // Format the date as you wish. Here's an example:
                        const year = date.getUTCFullYear();
                        const month = String(date.getUTCMonth() + 1).padStart(2, '0');  // Months are 0-indexed, so +1 is necessary
                        const day = String(date.getUTCDate()).padStart(2, '0');
                        const hours = String(date.getUTCHours()).padStart(2, '0');
                        const minutes = String(date.getUTCMinutes()).padStart(2, '0');

                        return `${year}-${month}-${day}`;
                    }
                },
                data: data.timestamp,

            },
            yAxis: {
                type: 'value',
                name: yAxisName
            },
            series: series.map(value => {
                return {
                    data: value.values.map((value, index) => [data.timestamp[index], value]),
                    type: "line",
                    name: value.dimension,
                    stack: "Total",
                    smooth: true
                }

            })
            // series: [{
            //     data: [],
            //     type: 'line',
            //     smooth: true
            // }]
        };
    }
    // window.addEventListener('resize', () => {
    //     chart.resize();
    // });

    useEffect(() => {
        const handleResize = () => {
            if (echartInstance.current) {
                echartInstance.current.getEchartsInstance().resize();
            }
        };

        window.addEventListener('resize', handleResize);
        setLoaded(true);
        return () => {
            window.removeEventListener('resize', handleResize);
        };
    }, []);


    return loaded ? (
        <div className="w-1/2 h-40" >
            <ReactEcharts ref={echartInstance} option={getOption()} />
        </div>

    ): null;
}

export default LineChart;
