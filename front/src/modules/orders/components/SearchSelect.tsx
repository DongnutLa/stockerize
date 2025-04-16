import { useState, useCallback } from 'react';
import { Select, TreeSelect } from 'antd';
import debounce from 'lodash/debounce';
import { BaseOptionType } from 'antd/es/select';
import { DataNode } from 'antd/es/tree';

interface DebouncedSearchSelectProps<T = any> {
    placeholder?: string
    valueProp: string
    fetchFn: (value: string) => Promise<T[]>
    options: BaseOptionType[] | DataNode[]
    setOptions: (opts: T[]) => void
    handleSelectOption: (value: string) => void
    excludedOptions: string[]
    isTreeSelect?: boolean
}

const DebouncedSearchSelect = <T,>({
    placeholder,
    fetchFn,
    valueProp,
    options,
    setOptions,
    handleSelectOption,
    excludedOptions,
    isTreeSelect,
}: DebouncedSearchSelectProps<T>) => {
  const [loading, setLoading] = useState(false);

  // Debounced API search function
  const fetchOptions = useCallback(
    debounce(async (searchValue) => {
      if (!searchValue) {
        setOptions([]);
        return;
      }

      setLoading(true);
      try {
        const response = await fetchFn(searchValue);
        let opts = [...response]
        if (excludedOptions.length) {
          opts = response.filter(res => excludedOptions.find(s => (res as any)[valueProp] ===  s))
        }
        setOptions(opts);
      } catch (error) {
        console.error("Error fetching options:", error);
      } finally {
        setLoading(false);
      }
    }, 700),
    [excludedOptions]
  );

  // Handle search input
  const handleSearch = (value: string) => {
    fetchOptions(value);
  };

  if (isTreeSelect) {
    return (
      <TreeSelect
        showSearch
        placeholder={placeholder ?? "Buscar..."}
        filterTreeNode={false}
        onSearch={handleSearch}
        loading={loading}
        style={{ width: "100%" }}
        onSelect={handleSelectOption}
        treeDefaultExpandAll
        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
        treeData={options as DataNode[]}
      />
    )
  }

  return (
    <Select
      showSearch
      placeholder={placeholder ?? "Buscar..."}
      filterOption={false} // Disable built-in filtering (let API handle it)
      onSearch={handleSearch}
      loading={loading}
      style={{ width: "100%" }}
      onSelect={handleSelectOption}
      options={options as BaseOptionType[]}
    />
  );
};

export default DebouncedSearchSelect;