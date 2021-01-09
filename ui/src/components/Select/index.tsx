import React from 'react';

import styles from './index.scss';

interface Props {
  onChange: (event: React.ChangeEvent<HTMLSelectElement>) => void;
  className?: string;
  id?: string;
  label?: string;
  value?: string | number | string[];
}

const Select: React.FC<Props> = (props) => (
  <div className={`${styles.selectContainer} ${props.className}`}>
    <label htmlFor={props.id} className={styles.label}>
      {props.label}
    </label>
    <select className={styles.select} onChange={props.onChange} value={props.value}>
      {props.children}
    </select>
    <span className={styles.arrow}></span>
  </div>
);

export default Select;