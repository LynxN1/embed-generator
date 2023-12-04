import { useState } from "react";
import { useScheduledMessageCreateMutation } from "../api/mutations";
import { useSendSettingsStore } from "../state/sendSettings";
import { useQueryClient } from "react-query";
import { useToasts } from "../util/toasts";
import EditorInput from "./EditorInput";
import Tooltip from "./Tooltip";
import {
  ArrowUpTrayIcon,
  XMarkIcon,
  ArrowRightIcon,
  CalendarDaysIcon,
  ArrowPathIcon,
} from "@heroicons/react/20/solid";
import clsx from "clsx";
import DateTimePicker from "./DateTimePicker";
import SavedMessageSelect from "./SavedMessageSelect";
import { ChannelSelect } from "./ChannelSelect";

export default function ScheduledMessageCreate({
  setCreate,
  cancelable,
}: {
  setCreate: (b: boolean) => void;
  cancelable: boolean;
}) {
  const guildId = useSendSettingsStore((s) => s.guildId);
  const createToast = useToasts((s) => s.create);

  const [name, setName] = useState("");
  const [onlyOnce, setOnlyOnce] = useState(true);
  const [startAt, setStartAt] = useState<string | undefined>();
  const [endAt, setEndAt] = useState<string | undefined>();
  const [cronExpression, setCronExpression] = useState("0 0 * * *");
  const [messageId, setMessageId] = useState<string | null>(null);
  const [channelId, setChannelId] = useState<string | null>(null);

  const queryClient = useQueryClient();
  const createMutation = useScheduledMessageCreateMutation();

  function create() {
    if (name.length == 0 || !guildId || !channelId || !messageId || !startAt) {
      createToast({
        title: "Some required fields are missing",
        message:
          "Please fill all the required fields before creating the scheduled message",
        type: "error",
      });
      return;
    }

    createMutation.mutate(
      {
        guildId: guildId,
        req: {
          name,
          description: null,
          channel_id: channelId,
          message_id: null,
          saved_message_id: messageId,
          cron_expression: cronExpression,
          start_at: startAt,
          end_at: endAt ?? null,
          only_once: onlyOnce,
          enabled: true,
        },
      },
      {
        onSuccess: (res) => {
          if (res.success) {
            setName("");
            setCreate(false);
            queryClient.invalidateQueries(["scheduled-messages", guildId]);
          } else {
            createToast({
              title: "Failed to create scheduled message",
              message: res.error.message,
              type: "error",
            });
          }
        },
      }
    );
  }

  return (
    <div className="bg-dark-3 p-5 rounded-lg">
      <div className="flex items-center space-x-2 text-lg mb-5 truncate justify-between">
        <div className="text-white truncate flex space-x-2 items-center">
          {onlyOnce ? (
            <CalendarDaysIcon className="text-gray-500 h-6 w-6" />
          ) : (
            <ArrowPathIcon className="text-gray-500 h-6 w-6" />
          )}
          <div>New Scheduled Message</div>
        </div>
        <div className="flex flex-none items-center space-x-4 md:space-x-3">
          {cancelable && (
            <div
              className="flex items-center text-gray-300 hover:text-white cursor-pointer md:bg-dark-2 md:rounded md:px-2 md:py-1"
              role="button"
              onClick={() => setCreate(false)}
            >
              <Tooltip text="Cancel">
                <XMarkIcon className="h-5 w-5" />
              </Tooltip>
              <div className="hidden md:block ml-2">Cancel</div>
            </div>
          )}
          <div
            className="flex items-center text-white cursor-pointer bg-blurple hover:bg-blurple-dark rounded px-2 py-1"
            role="button"
            onClick={create}
          >
            <Tooltip text="Create Scheduled Message">
              <ArrowUpTrayIcon className="h-5 w-5" />
            </Tooltip>
            <div className="ml-2">
              Create <span className="hidden md:inline-block">Schedule</span>
            </div>
          </div>
        </div>
      </div>
      <div className="space-y-5 mb-5">
        <EditorInput
          label="Name"
          type="text"
          maxLength={32}
          value={name}
          onChange={setName}
        />
        <div className="flex space-x-3 pb-3 items-end">
          <div className="flex-auto w-1/2">
            <div className="mb-1.5 flex">
              <div className="uppercase text-gray-300 text-sm font-medium">
                Message
              </div>
            </div>
            <SavedMessageSelect
              guildId={guildId}
              messageId={messageId}
              onChange={setMessageId}
            />
          </div>
          <div className="flex-none pb-2">
            <ArrowRightIcon className="h-5 w-5 text-gray-300" />
          </div>
          <div className="flex-auto w-1/2">
            <div className="mb-1.5 flex">
              <div className="uppercase text-gray-300 text-sm font-medium">
                Channel
              </div>
            </div>
            <ChannelSelect
              guildId={guildId}
              channelId={channelId}
              onChange={setChannelId}
            />
          </div>
        </div>
        <div className="flex">
          <button
            className="flex bg-dark-2 p-1 rounded text-white"
            onClick={() => setOnlyOnce((v) => !v)}
          >
            <div
              className={clsx(
                "py-1 px-2 rounded transition-colors",
                onlyOnce && "bg-dark-3"
              )}
            >
              Send Once
            </div>
            <div
              className={clsx(
                "py-1 px-2 rounded transition-colors",
                !onlyOnce && "bg-dark-3"
              )}
            >
              Send Periodically
            </div>
          </button>
        </div>
        {onlyOnce ? (
          <div>
            <div>
              <div className="mb-1.5 flex">
                <div className="uppercase text-gray-300 text-sm font-medium">
                  Send at
                </div>
              </div>
              <DateTimePicker
                value={startAt}
                onChange={setStartAt}
                clearable={false}
              />
            </div>
          </div>
        ) : (
          <>
            <div className="flex space-x-3">
              <div className="flex-auto">
                <div className="mb-1.5 flex">
                  <div className="uppercase text-gray-300 text-sm font-medium">
                    Start at
                  </div>
                </div>
                <DateTimePicker
                  value={startAt}
                  onChange={setStartAt}
                  clearable={false}
                />
              </div>
              <div className="flex-auto">
                <div className="mb-1.5 flex">
                  <div className="uppercase text-gray-300 text-sm font-medium">
                    End at
                  </div>
                </div>
                <DateTimePicker
                  value={endAt}
                  onChange={setEndAt}
                  clearable={true}
                />
              </div>
            </div>
            <div>
              <EditorInput
                label="CRON Expression"
                type="text"
                value={cronExpression}
                onChange={setCronExpression}
              />
            </div>
          </>
        )}
      </div>
    </div>
  );
}
