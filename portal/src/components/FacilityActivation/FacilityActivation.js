import React, { useState, useEffect } from "react";
import DropDown from "../DropDown/DropDown";
import { PROGRAMS, STATE_NAMES, DISTRICT_NAMES } from "../../utils/constants";
import styles from "./FacilityActivation.module.css";

function FacilityActivation() {
    const [listOfStates, setListOfStates] = useState([]);
    const [selectedProgram, setSelectedProgram] = useState();
    const [selectedState, setSelectedState] = useState("Karnataka");
    const [selectedDistrict, setSelectedDistrict] = useState("Bagalkote");
    const [facilityType, setFacilityType] = useState("Government");
    const [status, setStatus] = useState("Active");

    useEffect(() => {
        normalizeStateNames();
    }, []);

    const normalizeStateNames = () => {
        let data = [];
        Object.keys(STATE_NAMES).map((state) => {
            let newData = {};
            newData.value = state;
            newData.label = STATE_NAMES[state];
            data.push(newData);
        });
        setListOfStates(data);
    };

    const handleChange = (value, setValue) => {
        setValue(value);
        console.log(value);
    };

    const showDistrictList = () => {
        return Object.keys(DISTRICT_NAMES).map((district) => {
            return (
                <tr>
                    <td className={styles['highlight']}>
                        <div className="form-check">
                        <label className="form-check-label" htmlFor={district}>
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id={district}
                                name={district}
                                value={district}
                                onChange={(event) =>
                                    handleChange(district, setSelectedDistrict)
                                }
                                checked={selectedDistrict && district === selectedDistrict}
                            />
                            <div className={styles['wrapper']} style={{backgroundColor: selectedDistrict=== district ?'#DE9D00':''}}>&nbsp;</div>
                             {district}
                        </label>
                    </div>
                    </td>
                    <td>{DISTRICT_NAMES[district]}</td>
                </tr>
            );
        });
    };

    const handleClick = () => {
        console.log("clicked me")
    }
    return (
        <div className={`row ${styles['container']}`} >
            <div className="col-sm-3">
                <div>
                    <DropDown
                        options={PROGRAMS}
                        placeholder="Select Program"
                        setSelectedOption={setSelectedProgram}
                    />
                </div>
                <div>
                    <p className={styles['highlight']}>All of India</p>
                    <DropDown
                        options={listOfStates}
                        placeholder="Please select State"
                        setSelectedOption={setSelectedState}
                    />
                </div>
                <p className={styles['highlight']}>{selectedState}</p>
                <div className={`table-responsive ${styles["table"]}` }>
                    <table className="table table-borderless">
                        <thead>
                            <tr >Please select District</tr>
                        </thead>
                        <tbody className={styles['tbody']}>{showDistrictList()}</tbody>
                    </table>
                </div>
                <div>
                    <p className={styles['highlight']}>Type of Facility</p>
                    <div className="form-check">
                        <label className={`${'form-check-label'} ${styles['highlight']}`} htmlFor="government">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="government"
                                name="Government"
                                value="Government"
                                onClick={(event) =>
                                    handleChange(
                                        event.target.name,
                                        setFacilityType
                                    )
                                }
                                checked={facilityType === "Government"}
                            />
                            <div className={styles['wrapper']} style={{backgroundColor:facilityType==="Government"?'#DE9D00':''}}>&nbsp;</div>
                            Government
                        </label>
                    </div>
                    <div className="form-check">
                        <label className={`${'form-check-label'} ${styles['highlight']}`} htmlFor="private">
                            <input
                                type="checkbox"
                                className="form-check-input"
                                id="private"
                                name="Private"
                                value="Private"
                                onClick={(event) =>
                                    handleChange(
                                        event.target.name,
                                        setFacilityType
                                    )
                                }
                                checked={facilityType === "Private"}
                            />
                            <div className={styles['wrapper']} style={{backgroundColor:facilityType==="Private"?'#DE9D00':''}}>&nbsp;</div>
                            Private
                        </label>
                    </div>
                </div>
                <div>
                    <p className={styles['highlight']}>Status</p>
                    <div className="form-check">
                        <label className={`${'form-check-label'} ${styles['highlight']}`} htmlFor="Active">
                            <input
                                type="radio"
                                className="form-check-input"
                                id="Active"
                                name="Active"
                                value="Active"
                                onClick={(event) =>
                                    handleChange(
                                        event.target.name,
                                        setStatus
                                    )
                                }
                                checked={status === "Active"}
                            />
                            <div className={`${styles['wrapper']} ${styles['radio']}`}  style={{backgroundColor:status==="Active"?'#DE9D00':''}}>&nbsp;</div>
                            Active
                        </label>
                    </div>
                    <div className="form-check">
                        <label className={`${'form-check-label'} ${styles['highlight']}`} htmlFor="Inactive">
                            <input
                                type="radio"
                                className="form-check-input"
                                id="Inactive"
                                name="Inactive"
                                value="Inactive"
                                onClick={(event) =>
                                    handleChange(
                                        event.target.name,
                                        setStatus
                                    )
                                }
                                checked={status === "Inactive"}
                            />
                            <div className={`${styles['wrapper']} ${styles['radio']}`} style={{backgroundColor:status==="Inactive"?'#DE9D00':''}}>&nbsp;</div>
                            Inactive
                        </label>
                    </div>
                </div>
            </div>
            <div className="col-sm-6 container">
                <table className="table">
                <thead>
                    <tr>
                        <th>CENTRE ID</th>
                        <th>CENTRE NAME</th>
                        <th>VACCINATION STATIONS</th>
                        <th>CERTIFIED VACCINATORS</th>
                        <th>C19 program STATUS</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>100</td>
                        <td>This is a centre name</td>
                        <td>100</td>
                        <td>100</td>
                        <td>Inactive</td>
                    </tr>
                    <tr>
                        <td>100</td>
                        <td>This is a centre name</td>
                        <td>100</td>
                        <td>100</td>
                        <td>Inactive</td>
                    </tr><tr>
                        <td>100</td>
                        <td>This is a centre name</td>
                        <td>100</td>
                        <td>100</td>
                        <td>Inactive</td>
                    </tr>
                </tbody>
                </table>
            </div>
            <div className="col-sm-3 container">
                <div className="card">
                    <div className="card-body text-center">
                        <p>Make x facilities active for the x-program</p>
                        <button onClick={handleClick} className={styles['button']}>MAKE ACTIVE</button>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default FacilityActivation;