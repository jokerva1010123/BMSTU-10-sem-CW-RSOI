import "../Tables.css";
import { IStatistics } from "../../../interfaces/Statistics/IStatistics";


interface StatisticsRowProps {
	statistics: IStatistics;
	addClassName?: string;
}


export function StatisticsRow({ statistics, addClassName }: StatisticsRowProps) {
	return (
		<div 
			className={ `row ${ addClassName }` }
		>
			<div className="row-item basis-1/6">{ statistics.method }</div>
			<div className="row-item basis-1/2">{ statistics.url }</div>
			<div className="row-item basis-1/6">{ statistics.status_code }</div>
			<div className="row-item basis-1/4">{ statistics.time }</div>
		</div>
	)
}
